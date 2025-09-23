package handlers

import (
	"net/http"
	"rest-api/internal/models"
	commonutils "rest-api/internal/shared/common_utils"
	"rest-api/internal/shared/constants"
	userUseCase "rest-api/internal/user/usecases"

	"github.com/gin-gonic/gin"
	"github.com/remiges-tech/logharbour/logharbour"
)

type UserHandler struct {
	userUseCase userUseCase.UserUseCasesI
	logger      *logharbour.Logger
}

// CreateUserHandler implements UserHandlerI.
func (u *UserHandler) CreateUserHandler(ctx *gin.Context) {
	var (
		user   models.User
		logger = u.logger.WithModule("CreateUserHandler").WithPriority(logharbour.Info).WithModule("USERS")
	)
	//bind the request body to the user struct

	logger.Log("started create user handler")

	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Err().Error(err).Log("error binding json")
		field := "User"
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_ERROR_BINDING_JSON, constants.CODE_ERROR_BINDING_JSON, field, err.Error()),
		})
		return
	}
	if errorResponses := u.userUseCase.CreateUser(ctx, &user); len(errorResponses) > 0 {
		logger.Err().LogActivity("error creating user", map[string]any{
			"user_email": user.Email,
			"error":      errorResponses,
		})

		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, errorResponses)
		return
	}
	logger.Log("finished create user handler")

	commonutils.SendResponse(ctx, http.StatusCreated, user, nil)
}

// DeleteUserHandler implements UserHandlerI.
func (u *UserHandler) DeleteUserHandler(ctx *gin.Context) {
	var (
		idRequest models.IdRequest
		// add module + priority for clarity
		logger = u.logger.
			WithModule("DeleteUserHandler").
			WithPriority(logharbour.Info).
			WithModule("USERS")
	)

	logger.Log("started delete user handler")

	// Bind the request body
	if err := ctx.ShouldBindJSON(&idRequest); err != nil {
		logger.Err().
			Error(err).
			Log("error binding json to IdRequest")

		field := "User"
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil,
			[]commonutils.ErrorResponse{
				commonutils.BuildErrorResponse(
					constants.MSG_ERROR_BINDING_JSON,
					constants.CODE_ERROR_BINDING_JSON,
					field,
					err.Error(),
				),
			})
		return
	}

	logger.LogActivity("parsed request", map[string]any{
		"user_id": idRequest.ID,
	})

	// Call use-case layer
	if errorResponses := u.userUseCase.DeleteUser(ctx, idRequest.ID); len(errorResponses) > 0 {
		logger.Err().
			LogActivity("error deleting user", map[string]any{
				"user_id": idRequest.ID,
				"errors":  errorResponses,
			})

		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, errorResponses)
		return
	}

	logger.LogActivity("user deleted successfully", map[string]any{
		"user_id": idRequest.ID,
	})

	logger.Log("finished delete user handler")

	commonutils.SendResponse(ctx, http.StatusOK, nil, nil)
}

// GetAllUsersHandler implements UserHandlerI.
func (u *UserHandler) GetAllUsersHandler(ctx *gin.Context) {
	var (
		listRequest models.ListRequest
		logger      = u.logger.
				WithModule("GetAllUsersHandler").
				WithPriority(logharbour.Info).
				WithModule("USERS")
	)

	logger.Log("started get all users handler")

	// Bind request body
	if err := ctx.ShouldBindJSON(&listRequest); err != nil {
		logger.Err().
			Error(err).
			Log("error binding json to ListRequest")

		field := "User"
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil,
			[]commonutils.ErrorResponse{
				commonutils.BuildErrorResponse(
					constants.MSG_ERROR_BINDING_JSON,
					constants.CODE_ERROR_BINDING_JSON,
					field,
					err.Error(),
				),
			})
		return
	}

	logger.LogActivity("parsed request", map[string]any{
		"offset": listRequest.FilterParams.OffSet,
		"limit":  listRequest.FilterParams.Limit,
	})

	// Call use-case layer
	users, errorResponses := u.userUseCase.GetAllUsers(ctx, listRequest.FilterParams)
	if len(errorResponses) > 0 {
		logger.Err().
			LogActivity("error fetching users", map[string]any{
				"offset": listRequest.FilterParams.OffSet,
				"limit":  listRequest.FilterParams.Limit,
				"errors": errorResponses,
			})

		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, errorResponses)
		return
	}

	logger.LogActivity("fetched users successfully", map[string]any{
		"count": len(users),
	})

	logger.Log("finished get all users handler")

	commonutils.SendResponse(ctx, http.StatusOK, users, nil)
}

// GetUserHandler implements UserHandlerI.
func (u *UserHandler) GetUserHandler(ctx *gin.Context) {
	var (
		idRequest models.IdRequest
		logger    = u.logger.
				WithModule("GetUserHandler").
				WithPriority(logharbour.Info).
				WithModule("USERS")
	)

	logger.Log("started get user handler")

	// Bind the request body to IdRequest
	if err := ctx.ShouldBindJSON(&idRequest); err != nil {
		logger.Err().
			Error(err).
			Log("error binding json to IdRequest")

		field := "User"
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil,
			[]commonutils.ErrorResponse{
				commonutils.BuildErrorResponse(
					constants.MSG_ERROR_BINDING_JSON,
					constants.CODE_ERROR_BINDING_JSON,
					field,
					err.Error(),
				),
			})
		return
	}

	logger.LogActivity("parsed request", map[string]any{
		"user_id": idRequest.ID,
	})

	// Fetch the user
	user, errorResponses := u.userUseCase.GetUser(ctx, idRequest.ID)
	if len(errorResponses) > 0 {
		logger.Err().
			LogActivity("error fetching user", map[string]any{
				"user_id": idRequest.ID,
				"errors":  errorResponses,
			})

		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, errorResponses)
		return
	}

	logger.LogActivity("fetched user successfully", map[string]any{
		"user_id": user.Id,
	})

	logger.Log("finished get user handler")

	commonutils.SendResponse(ctx, http.StatusOK, user, nil)
}

// UpdateUserHandler implements UserHandlerI.
func (u *UserHandler) UpdateUserHandler(ctx *gin.Context) {
	var (
		user   models.User
		logger = u.logger.
			WithModule("UpdateUserHandler").
			WithPriority(logharbour.Info).
			WithModule("USERS")
	)

	logger.Log("started update user handler")

	// Bind the request body to User struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Err().
			Error(err).
			Log("error binding json to User")

		field := "User"
		commonutils.SendResponse(ctx, http.StatusBadRequest, nil,
			[]commonutils.ErrorResponse{
				commonutils.BuildErrorResponse(
					constants.MSG_ERROR_BINDING_JSON,
					constants.CODE_ERROR_BINDING_JSON,
					field,
					err.Error(),
				),
			})
		return
	}

	logger.LogActivity("parsed request", map[string]any{
		"user_id":    user.Id,
		"user_email": user.Email,
	})

	// Call the use-case layer
	if errorResponses := u.userUseCase.UpdateUser(ctx, user); len(errorResponses) > 0 {
		logger.Err().
			LogActivity("error updating user", map[string]any{
				"user_id": user.Id,
				"errors":  errorResponses,
			})

		commonutils.SendResponse(ctx, http.StatusBadRequest, nil, errorResponses)
		return
	}

	logger.LogActivity("user updated successfully", map[string]any{
		"user_id": user.Id,
	})

	logger.Log("finished update user handler")

	commonutils.SendResponse(ctx, http.StatusOK, nil, nil)
}

type UserHandlerI interface {
	CreateUserHandler(ctx *gin.Context)
	GetAllUsersHandler(ctx *gin.Context)
	GetUserHandler(ctx *gin.Context)
	UpdateUserHandler(ctx *gin.Context)
	DeleteUserHandler(ctx *gin.Context)
}

func NewUserHandler(userUseCase userUseCase.UserUseCasesI,logger *logharbour.Logger) UserHandlerI {
	return &UserHandler{
		userUseCase: userUseCase,
		logger:      logger,
	}
}

func RegisterRoutes(r *gin.Engine, userHandlerI UserHandlerI) {

	user := r.Group("api/v1/student/")
	{
		user.POST("/create", userHandlerI.CreateUserHandler)
		user.GET("/list", userHandlerI.GetAllUsersHandler)
		user.GET("/get", userHandlerI.GetUserHandler)
		user.PUT("/update", userHandlerI.UpdateUserHandler)
		user.PUT("/delete", userHandlerI.DeleteUserHandler)
	}

}
