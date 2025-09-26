package repo

import (
	"context"
	"fmt"
	"rest-api/db/sqlc"
	"rest-api/internal/models"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepo struct {
	querier sqlc.Querier
}

type UserRepoI interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetAllUsers(ctx context.Context, params models.FilterParams) ([]models.User, error)
	GetUser(ctx context.Context, id int64) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id int64) error
}

func NewUserRepo(querier sqlc.Querier) UserRepoI {

	return &UserRepo{
		querier: querier,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user *models.User) error {
	id, err := u.querier.CreateUser(ctx, sqlc.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		PhoneNumber: pgtype.Text{
			String: user.PhoneNumber,
			Valid:  user.PhoneNumber != "",
		},
	})
	if err != nil {
		return err
	}
	user.Id = int64(id)
	return nil

}
func (u *UserRepo) GetAllUsers(ctx context.Context, params models.FilterParams) ([]models.User, error) {
	sqlcUsers, err := u.querier.ListUsers(ctx, sqlc.ListUsersParams{
		Offset: int32(params.OffSet),
		Limit:  int32(params.Limit),
	})
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}
	return sqlcUsersToModels(sqlcUsers), nil
}

func sqlcUsersToModels(users []sqlc.User) []models.User {
	out := make([]models.User, 0, len(users))
	for _, usr := range users {
		out = append(out, models.User{
			Id:          int64(usr.ID),
			FirstName:   usr.FirstName,
			LastName:    usr.LastName,
			Email:       usr.Email,
			PhoneNumber: usr.PhoneNumber.String,
		})
	}
	return out
}

func (u *UserRepo) GetUser(ctx context.Context, id int64) (models.User, error) {
	user, err := u.querier.GetUser(ctx, int32(id))
	if err != nil {
		return models.User{}, fmt.Errorf("get user: %w", err)
	}
	return models.User{
		Id:          int64(user.ID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber.String,
	}, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, user models.User) error {
	err := u.querier.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:        int32(user.Id),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		PhoneNumber: pgtype.Text{
			String: user.PhoneNumber,
			Valid:  user.PhoneNumber != "",
		},
	})
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}
	return nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, id int64) error {
	err := u.querier.DeleteUser(ctx, int32(id))
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}
