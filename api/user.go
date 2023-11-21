package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/patrickchap/clipsapi/db/sqlc"
)

type AddUserParams struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Auth0UserID string `json:"auth0_user_id" binding:"required"`
}

type UserResponse struct {
	Username    string `form:"username"`
	Email       string `form:"email"`
}
func (server *Server) addUser(ctx *gin.Context){
	var req AddUserParams

	err := ctx.ShouldBindJSON(&req); if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Email: req.Email,
		Auth0UserID: req.Auth0UserID,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userResponse := UserResponse{
		Username: user.Username,
		Email: user.Email,
	}

	ctx.JSON(http.StatusOK, userResponse)
}
