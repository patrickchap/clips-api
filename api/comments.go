package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/patrickchap/clipsapi/db/sqlc"
)

type getVideoCommentsReq struct {
	VideoID int64 `uri:"video_id" binding:"required"`
}
type getVideoCommentsForm struct {
	Limit	int64 `form:"limit" binding:"required"`
	Offset	int64 `form:"offset"`
}
func (server *Server) getVideoComments(ctx *gin.Context){
	var req getVideoCommentsReq
	var reqForm getVideoCommentsForm

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	if err := ctx.ShouldBindQuery(&reqForm); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	arg := db.GetCommentsByVideoParams{
		VideoID: req.VideoID,
		Limit: reqForm.Limit,
		Offset: (reqForm.Limit - 1) * reqForm.Offset,
	}

	comments, err := server.store.GetCommentsByVideo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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

func (server *Server) addVideoComments(ctx *gin.Context){
	var req addVideoCommentsReq
	var reqForm addVideoCommentsReqForm

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	if err := ctx.ShouldBindJSON(&reqForm); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	video, err := server.store.GetVideo(ctx, req.VideoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}


	if !validateClaims(ctx, video.UserID){
		return
	}

	arg := db.CreateCommentParams{
		VideoID: req.VideoID,
		UserID: video.UserID,
		Content: reqForm.Content,
	}

	comments, err := server.store.CreateComment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

type deleteVideoCommentsReq struct {
	VideoID	   pgtype.Int8	`uri:"video_id" binding:"required"`
}

func (server *Server) deleteVideoComments(ctx *gin.Context){
	var req addVideoCommentsReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	if err := server.store.DeleteCommentsByVideo(ctx, req.VideoID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	message := fmt.Sprintf("Video with ID %v deleted", req.VideoID)
	ctx.JSON(http.StatusOK, gin.H{"message": message})

}

