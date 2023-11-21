package api

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	db "github.com/patrickchap/clipsapi/db/sqlc"
)

type getVideoListParams struct {
	Limit  int64 `form:"limit" binding:"required,min=1"`
	Offset int64 `form:"offset"`
	Search string `form:"search"`
}

func (server *Server) getListVideo(ctx *gin.Context){
	var req getVideoListParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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

func (server *Server) getUserVideoList(ctx *gin.Context){
	var req getUserVideoListUriIdParams
	var reqForm getUserVideoListFormParams 

	// Bind URI parameters
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Bind query parameters
	if err := ctx.ShouldBindQuery(&reqForm); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println(req)
	arg := db.GetUserVideoWithLikesParams{
		UserID: req.UserID,
		Limit: reqForm.Limit,
		Offset: (reqForm.Limit-1)*reqForm.Offset,
	}

	fmt.Println(arg)

	videos, err := server.store.GetUserVideoWithLikes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}


	ctx.JSON(http.StatusOK, videos)
}



type getVideoParams struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getVideo(ctx *gin.Context){
	var req getVideoParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	
	video, err := server.store.GetVideoWithLikes(ctx, req.ID)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, video)

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

func (server *Server) createVideo(ctx *gin.Context){
	var req createVideoParams

	const (
		s3BucketName = "clips-bucket-prch"
		s3Region = "us-west-2"
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}


	videoResultsChan := make(chan uploadResult)
	thumbnailResultsChan := make(chan uploadResult)

	defer close(videoResultsChan)
	defer close(thumbnailResultsChan)

	go func() {
		result, err := uploadFileToS3(req.File, s3BucketName, s3Region)
		if err != nil {
		    videoResultsChan <- uploadResult{err: err}
		    return
		}
		videoResultsChan <- uploadResult{result: result}
	}()

	go func() {
		result, err := uploadFileToS3(req.Thumbnail, s3BucketName, s3Region)
		if err != nil {
		    thumbnailResultsChan <- uploadResult{err: err}
		    return
		}
		thumbnailResultsChan <- uploadResult{result: result}
	}()

	videoResult := <-videoResultsChan
	thumbnailResult := <-thumbnailResultsChan


	if videoResult.err != nil || thumbnailResult.err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("Error uploading to s3")))
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, video)
	
}


func uploadFileToS3(req *multipart.FileHeader, bucketName string, region string) (*s3manager.UploadOutput, error) {
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



