package service

import (
	"context"
	"rest-api/internal/models"
	commonutils "rest-api/internal/shared/common_utils"
	"rest-api/internal/shared/constants"
	userRepo "rest-api/internal/user/repo"
)

type UserServiceI interface {
	CreateUser(ctx context.Context, user *models.User) []commonutils.ErrorResponse
	GetAllUsers(ctx context.Context, params models.FilterParams) ([]models.User, []commonutils.ErrorResponse)
	GetUser(ctx context.Context, id int64) (models.User, []commonutils.ErrorResponse)
	UpdateUser(ctx context.Context, user models.User) []commonutils.ErrorResponse
	DeleteUser(ctx context.Context, id int64) []commonutils.ErrorResponse
}

type UserService struct {
	repo userRepo.UserRepoI
}

// CreateUser implements UserServiceI.
func (u *UserService) CreateUser(ctx context.Context, user *models.User) []commonutils.ErrorResponse {

	//any additional user validation / logic can be added here

	err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_DATABASE_ERROR, constants.CODE_DATABASE_ERROR, "User", err.Error()),
		}
	}

	return nil
}

// DeleteUser implements UserServiceI.
func (u *UserService) DeleteUser(ctx context.Context, id int64) []commonutils.ErrorResponse {
	//any additional user validation / logic can be added here

	err := u.repo.DeleteUser(ctx, id)
	if err != nil {
		return []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_DATABASE_ERROR, constants.CODE_DATABASE_ERROR, "User", err.Error()),
		}
	}
	return nil
}

// GetAllUsers implements UserServiceI.
func (u *UserService) GetAllUsers(ctx context.Context, params models.FilterParams) ([]models.User, []commonutils.ErrorResponse) {
	//any additional user validation / logic can be added here

	users, err := u.repo.GetAllUsers(ctx, params)
	if err != nil {
		return nil, []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_DATABASE_ERROR, constants.CODE_DATABASE_ERROR, "User", err.Error()),
		}
	}
	return users, nil
}

// GetUser implements UserServiceI.
func (u *UserService) GetUser(ctx context.Context, id int64) (models.User, []commonutils.ErrorResponse) {
	//any additional user validation / logic can be added here

	user, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return models.User{}, []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_DATABASE_ERROR, constants.CODE_DATABASE_ERROR, "User", err.Error()),
		}
	}
	return user, nil
}

// UpdateUser implements UserServiceI.
func (u *UserService) UpdateUser(ctx context.Context, user models.User) []commonutils.ErrorResponse {
	//any additional user validation / logic can be added here

	err := u.repo.UpdateUser(ctx, user)
	if err != nil {
		return []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_DATABASE_ERROR, constants.CODE_DATABASE_ERROR, "User", err.Error()),
		}
	}
	return nil
}

func NewUserService(repo userRepo.UserRepoI) UserServiceI {
	return &UserService{
		repo: repo,
	}
}
