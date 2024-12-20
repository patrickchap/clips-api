package controllers

import (
	"fmt"
	"mime/multipart"

	"net/http"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/patrickchap/clipsapi/util"
	db "github.com/patrickchap/clipsapi/db/sqlc"
)

type VideoController struct {
	store db.Store
}

func NewVideoController(s db.Store) *VideoController {
    return &VideoController{
        store: s,
    }
}

type getVideoListParams struct {
	Limit  int64 `form:"limit" binding:"required,min=1"`
	Offset int64 `form:"offset"`
	Search string `form:"search"`
}

func (server *VideoController) GetListVideo(ctx *gin.Context){
	var req getVideoListParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	wildcard := "%"
	// %% to search for all
	search := wildcard + wildcard 
	
	if req.Search != "" {
		search =  wildcard + req.Search + wildcard
	}

	arg := db.ListVideosWithLikesAndSearchParams{
		Offset: (req.Limit-1)*req.Offset,
		Limit: req.Limit,
		Search: search,
	}

	
	video, err := server.store.ListVideosWithLikesAndSearch(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, video)

}

type getUserVideoListFormParams struct {
	Limit  int64  `form:"limit" binding:"required,min=1"`
	Offset int64  `form:"offset"`
}

type getUserVideoListUriIdParams struct {
    UserID string `uri:"user_id" binding:"required"`
}

func (server *VideoController) GetUserVideoList(ctx *gin.Context){
	var req getUserVideoListUriIdParams
	var reqForm getUserVideoListFormParams 

	// Bind URI parameters
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	// Bind query parameters
	if err := ctx.ShouldBindQuery(&reqForm); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := db.GetUserVideoWithLikesParams{
		UserID: req.UserID,
		Limit: reqForm.Limit,
		Offset: (reqForm.Limit-1)*reqForm.Offset,
	}

	videos, err := server.store.GetUserVideoWithLikes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, videos)
}



type getVideoParams struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *VideoController) GetVideo(ctx *gin.Context){
	var req getVideoParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	
	video, err := server.store.GetVideoWithLikes(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, video)

}

type updateVideoUriParams struct {
    Id int64 `uri:"id" binding:"required"`
}

type updateVideoFormParams struct {
    Title string `form:"title"`
    Description string `form:"description"`
}

func (server *VideoController) UpdateVideo(ctx *gin.Context){
	var reqUri updateVideoUriParams
	var reqForm updateVideoFormParams

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqForm); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	video, err := server.store.GetVideo(ctx, reqUri.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if !util.ValidateClaims(ctx, video.UserID){
		return
	}

	arg := db.UpdateVideoParams{
		ID: video.ID,
		Title: reqForm.Title,
		Description: reqForm.Description,
	}


	if err := server.store.UpdateVideo(ctx, arg); err != nil {

		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Video updated successfully"})
}

// Define a struct to hold the result of an upload operation
type uploadResult struct {
    result *s3manager.UploadOutput
    err    error
}

type createVideoParams struct {
	Title       string	           `form:"title" binding:"required"`
	Description string	           `form:"description" binding:"required"`
	File        *multipart.FileHeader  `form:"file" binding:"required"`
	Thumbnail   *multipart.FileHeader  `form:"thumbnail" binding:"required"`
	UserID      string                  `form:"userId" binding:"required"`
}

func (server *VideoController) CreateVideo(ctx *gin.Context){
	var req createVideoParams

	const (
		s3BucketName = "clips-bucket-prch"
		s3Region = "us-west-2"
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}


	videoResultsChan := make(chan uploadResult)
	thumbnailResultsChan := make(chan uploadResult)

	defer close(videoResultsChan)
	defer close(thumbnailResultsChan)

	go func() {
		result, err := UploadFileToS3(req.File, s3BucketName, s3Region)
		if err != nil {
		    videoResultsChan <- uploadResult{err: err}
		    return
		}
		videoResultsChan <- uploadResult{result: result}
	}()

	go func() {
		result, err := UploadFileToS3(req.Thumbnail, s3BucketName, s3Region)
		if err != nil {
		    thumbnailResultsChan <- uploadResult{err: err}
		    return
		}
		thumbnailResultsChan <- uploadResult{result: result}
	}()

	videoResult := <-videoResultsChan
	thumbnailResult := <-thumbnailResultsChan


	if videoResult.err != nil || thumbnailResult.err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(fmt.Errorf("Error uploading to s3")))
		return
	}


	fileUrl := videoResult.result.Location
	thumbnailUrl := thumbnailResult.result.Location 

	
	arg := db.CreateVideoParams{
		Title: req.Title,
		Description: req.Description,
		FileUrl: fileUrl,
		UserID: req.UserID,
		ThumbnailUrl: thumbnailUrl,
	}

	video, err := server.store.CreateVideo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, video)
	
}

func UploadFileToS3Async(file *multipart.FileHeader,  resultsChan chan <- uploadResult){
	const (
		s3BucketName = "clips-bucket-prch"
		s3Region = "us-west-2"
	)
	result, err := UploadFileToS3(file, s3BucketName, s3Region)
	resultsChan <- uploadResult{result: result, err: err}
}

func UploadFileToS3(req *multipart.FileHeader, bucketName string, region string) (*s3manager.UploadOutput, error) {
	file, err := req.Open()
	if err != nil {
		return nil, fmt.Errorf("Open failed: %v", err)
	}
	defer file.Close()

	session, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to create AWS session: %v", err)
	}

	uploader := s3manager.NewUploader(session)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(req.Filename),
		Body:   file,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return nil, fmt.Errorf("Save to S3 failed: %v", err)
	}

	return result, nil 
}




