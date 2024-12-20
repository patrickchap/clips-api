package controllers 

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/patrickchap/clipsapi/db/sqlc"
	"github.com/patrickchap/clipsapi/util"
)

type CommentController struct {
	store db.Store
}

func NewCommentController(s db.Store) *CommentController {
    return &CommentController{
        store: s,
    }
}
type getVideoCommentsReq struct {
	VideoID int64 `uri:"video_id" binding:"required"`
}
type getVideoCommentsForm struct {
	Limit	int64 `form:"limit" binding:"required"`
	Offset	int64 `form:"offset"`
}
func (comment *CommentController) GetVideoComments(ctx *gin.Context){
	var req getVideoCommentsReq
	var reqForm getVideoCommentsForm

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return 
	}

	if err := ctx.ShouldBindQuery(&reqForm); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return 
	}

	arg := db.GetCommentsByVideoParams{
		VideoID: req.VideoID,
		Limit: reqForm.Limit,
		Offset: (reqForm.Limit - 1) * reqForm.Offset,
	}

	comments, err := comment.store.GetCommentsByVideo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

type addVideoCommentsReq struct {
	VideoID	int64 `uri:"video_id" binding:"required"`
}

type addVideoCommentsReqForm struct {
	Content string `form:"video_id" binding:"required"`
}

func (comment *CommentController) AddVideoComments(ctx *gin.Context){
	var req addVideoCommentsReq
	var reqForm addVideoCommentsReqForm

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return 
	}

	if err := ctx.ShouldBindJSON(&reqForm); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return 
	}

	video, err := comment.store.GetVideo(ctx, req.VideoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}


	if !util.ValidateClaims(ctx, video.UserID){
		return
	}

	arg := db.CreateCommentParams{
		VideoID: req.VideoID,
		UserID: video.UserID,
		Content: reqForm.Content,
	}

	comments, err := comment.store.CreateComment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

type deleteVideoCommentsReq struct {
	VideoID int64 `uri:"video_id" binding:"required"`
}

func (comment *CommentController) DeleteVideoComments(ctx *gin.Context){
	var req addVideoCommentsReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return 
	}

	if err := comment.store.DeleteCommentsByVideo(ctx, req.VideoID); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return 
	}

	message := fmt.Sprintf("Video with ID %v deleted", req.VideoID)
	ctx.JSON(http.StatusOK, gin.H{"message": message})

}

