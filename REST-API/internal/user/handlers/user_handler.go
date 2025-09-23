package handlers

import (
	"net/http"
	"rest-api/internal/models"
	commonutils "rest-api/internal/shared/common_utils"
	"rest-api/internal/shared/constants"
	userUseCase "rest-api/internal/user/usecases"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase userUseCase.UserUseCasesI
}

// CreateUserHandler implements UserHandlerI.
func (u *UserHandler) CreateUserHandler(ctx *gin.Context) {
	var (
		user models.User
	)
	//bind the request body to the user struct

	if err := ctx.ShouldBindJSON(&user); err != nil {
		field := "User"
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_ERROR_BINDING_JSON, constants.CODE_ERROR_BINDING_JSON, field, err.Error()),
		})
		return
	}
	if errorResponses := u.userUseCase.CreateUser(ctx, &user); errorResponses != nil {
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, errorResponses)
		return
	}

	commonutils.SendResponse(ctx, http.StatusOK, user, nil)
}

// DeleteUserHandler implements UserHandlerI.
func (u *UserHandler) DeleteUserHandler(ctx *gin.Context) {
	var (
		idRequest models.IdRequest
	)
	if err := ctx.ShouldBindJSON(&idRequest); err != nil {
		field := "User"
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_ERROR_BINDING_JSON, constants.CODE_ERROR_BINDING_JSON, field, err.Error()),
		})
		return
	}
	if errorResponses := u.userUseCase.DeleteUser(ctx, idRequest.ID); errorResponses != nil {
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, errorResponses)
		return
	}
	commonutils.SendResponse(ctx, http.StatusOK, nil, nil)

}

// GetAllUsersHandler implements UserHandlerI.
func (u *UserHandler) GetAllUsersHandler(ctx *gin.Context) {
	var (
		listRequest models.ListRequest
	)

	if err := ctx.ShouldBindJSON(&listRequest); err != nil {
		field := "User"
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_ERROR_BINDING_JSON, constants.CODE_ERROR_BINDING_JSON, field, err.Error()),
		})
		return
	}
	users, errorResponses := u.userUseCase.GetAllUsers(ctx, listRequest.FilterParams)
	if errorResponses != nil {
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, errorResponses)
		return
	}
	commonutils.SendResponse(ctx, http.StatusOK, users, nil)
}

// GetUserHandler implements UserHandlerI.
func (u *UserHandler) GetUserHandler(ctx *gin.Context) {
	var (
		idRequest models.IdRequest
	)

	if err:=ctx.ShouldBindJSON(&idRequest); err != nil {
		field := "User"
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_ERROR_BINDING_JSON, constants.CODE_ERROR_BINDING_JSON, field, err.Error()),
		})
		return
	}
	user, errorResponses := u.userUseCase.GetUser(ctx, idRequest.ID)
	if errorResponses != nil {
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, errorResponses)
		return
	}
	commonutils.SendResponse(ctx, http.StatusOK, user, nil)
}

// UpdateUserHandler implements UserHandlerI.
func (u *UserHandler) UpdateUserHandler(ctx *gin.Context) {
	var (
		user models.User
	)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		field := "User"
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_ERROR_BINDING_JSON, constants.CODE_ERROR_BINDING_JSON, field, err.Error()),
		})
		return
	}
	if errorResponses := u.userUseCase.UpdateUser(ctx, user); errorResponses != nil {
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, errorResponses)
		return
	}
	commonutils.SendResponse(ctx, http.StatusOK, nil, nil)
}

type UserHandlerI interface {
	CreateUserHandler(ctx *gin.Context)
	GetAllUsersHandler(ctx *gin.Context)
	GetUserHandler(ctx *gin.Context)
	UpdateUserHandler(ctx *gin.Context)
	DeleteUserHandler(ctx *gin.Context)
}

func NewUserHandler(userUseCase userUseCase.UserUseCasesI) UserHandlerI {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}
