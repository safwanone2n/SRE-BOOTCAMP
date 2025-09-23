package usecases

import (
	"context"
	"rest-api/internal/models"
	commonutils "rest-api/internal/shared/common_utils"
	"rest-api/internal/shared/constants"
	userRepo "rest-api/internal/user/repo"
)

type UserUseCasesI interface {
	CreateUser(ctx context.Context, user *models.User) []commonutils.ErrorResponse
	GetAllUsers(ctx context.Context, params models.FilterParams) ([]models.User, []commonutils.ErrorResponse)
	GetUser(ctx context.Context, id int64) (models.User, []commonutils.ErrorResponse)
	UpdateUser(ctx context.Context, user models.User) []commonutils.ErrorResponse
	DeleteUser(ctx context.Context, id int64) []commonutils.ErrorResponse
}

type UserUseCases struct {
	repo userRepo.UserRepoI
}

// CreateUser implements UserUseCasesI.
func (u *UserUseCases) CreateUser(ctx context.Context, user *models.User) []commonutils.ErrorResponse {

	//any additional user validation / logic can be added here

	err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_DATABASE_ERROR, constants.CODE_DATABASE_ERROR, "User", err.Error()),
		}
	}

	return nil
}

// DeleteUser implements UserUseCasesI.
func (u *UserUseCases) DeleteUser(ctx context.Context, id int64) []commonutils.ErrorResponse {
	//any additional user validation / logic can be added here

	err := u.repo.DeleteUser(ctx, id)
	if err != nil {
		return []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_DATABASE_ERROR, constants.CODE_DATABASE_ERROR, "User", err.Error()),
		}
	}
	return nil
}

// GetAllUsers implements UserUseCasesI.
func (u *UserUseCases) GetAllUsers(ctx context.Context, params models.FilterParams) ([]models.User, []commonutils.ErrorResponse) {
	//any additional user validation / logic can be added here

	users, err := u.repo.GetAllUsers(ctx, params)
	if err != nil {
		return nil, []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_DATABASE_ERROR, constants.CODE_DATABASE_ERROR, "User", err.Error()),
		}
	}
	return users, nil
}

// GetUser implements UserUseCasesI.
func (u *UserUseCases) GetUser(ctx context.Context, id int64) (models.User, []commonutils.ErrorResponse) {
	//any additional user validation / logic can be added here

	user, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return models.User{}, []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_DATABASE_ERROR, constants.CODE_DATABASE_ERROR, "User", err.Error()),
		}
	}
	return user, nil
}

// UpdateUser implements UserUseCasesI.
func (u *UserUseCases) UpdateUser(ctx context.Context, user models.User) []commonutils.ErrorResponse {
	//any additional user validation / logic can be added here

	err := u.repo.UpdateUser(ctx, user)
	if err != nil {
		return []commonutils.ErrorResponse{
			commonutils.BuildErrorResponse(constants.MSG_DATABASE_ERROR, constants.CODE_DATABASE_ERROR, "User", err.Error()),
		}
	}
	return nil
}

func NewUserUseCases(repo userRepo.UserRepoI) UserUseCasesI {
	return &UserUseCases{
		repo: repo,
	}
}
