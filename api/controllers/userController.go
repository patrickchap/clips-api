package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/patrickchap/clipsapi/db/sqlc"
	"github.com/patrickchap/clipsapi/util"
)

type UserController struct {
	store db.Store
}

func NewUserController(s db.Store) *UserController {
    return &UserController{
        store: s,
    }
}

type AddUserParams struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Auth0UserID string `json:"auth0_user_id" binding:"required"`
}

type UserResponse struct {
	Username    string `form:"username"`
	Email       string `form:"email"`
}

func (userController *UserController) AddUser(ctx *gin.Context){
	var req AddUserParams

	err := ctx.ShouldBindJSON(&req); if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Email: req.Email,
		Auth0UserID: req.Auth0UserID,
	}

	user, err := userController.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	userResponse := UserResponse{
		Username: user.Username,
		Email: user.Email,
	}

	ctx.JSON(http.StatusOK, userResponse)
}
