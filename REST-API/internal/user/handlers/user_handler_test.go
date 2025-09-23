package handlers

import (
	"net/http"
	"rest-api/internal/models"
	commonutils "rest-api/internal/shared/common_utils"
	testutils "rest-api/internal/shared/test_utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestCreateUserHandler(t *testing.T) {
	// Table of test cases
	logger := testutils.GetLogger()
	testCases := []struct {
		name               string
		body               any
		mockFunc           func(m *testutils.HandlerDeps)
		expectedStatusCode int
	}{
		{
			name: "success - create user",
			body: models.User{
				Id:          1,
				FirstName:   "test",
				LastName:    "test",
				Email:       "test",
				PhoneNumber: "test",
			},
			mockFunc: func(m *testutils.HandlerDeps) {
				m.UserUseCaseMock.
					EXPECT().
					CreateUser(gomock.Any(), gomock.AssignableToTypeOf(&models.User{})).
					Return([]commonutils.ErrorResponse{})
			},
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			deps := testutils.NewHandlerDeps(t)

			// set expectations
			tc.mockFunc(deps)
			w := testutils.DoRequest(t,
				func(r *gin.Engine) {
					h := NewUserHandler(deps.UserUseCaseMock, logger)
					r.POST("/users/create", h.CreateUserHandler)
				},
				http.MethodPost, "/users/create", tc.body)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("expected %d, got %d", tc.expectedStatusCode, w.Code)
			}
		})
	}
}

func TestGetAllUsersHandler(t *testing.T) {
	logger := testutils.GetLogger()

	// Table of test cases
	testCases := []struct {
		name               string
		body               any
		mockFunc           func(m *testutils.HandlerDeps)
		expectedStatusCode int
	}{
		{
			name: "success - get all users",
			body: models.ListRequest{
				FilterParams: models.FilterParams{
					OffSet: 0,
					Limit:  10,
				},
			},
			mockFunc: func(m *testutils.HandlerDeps) {
				m.UserUseCaseMock.
					EXPECT().
					GetAllUsers(gomock.Any(), gomock.AssignableToTypeOf(models.FilterParams{})).
					Return([]models.User{{
						Id:          1,
						FirstName:   "test",
						LastName:    "test",
						Email:       "test",
						PhoneNumber: "test",
					},
						{
							Id:          2,
							FirstName:   "test",
							LastName:    "test",
							Email:       "test",
							PhoneNumber: "test",
						},
					}, []commonutils.ErrorResponse{})
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			deps := testutils.NewHandlerDeps(t)

			// set expectations
			tc.mockFunc(deps)

			w := testutils.DoRequest(t,
				func(r *gin.Engine) {
					h := NewUserHandler(deps.UserUseCaseMock, logger)
					r.GET("/users/list", h.GetAllUsersHandler)
				},
				http.MethodGet, "/users/list", tc.body)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("expected %d, got %d", tc.expectedStatusCode, w.Code)
			}
		})
	}
}

func TestGetUserHandler(t *testing.T) {
	logger := testutils.GetLogger()

	// Table of test cases
	testCases := []struct {
		name               string
		body               any
		mockFunc           func(m *testutils.HandlerDeps)
		expectedStatusCode int
	}{
		{
			name: "success - get user",
			body: models.IdRequest{
				ID: 1,
			},
			mockFunc: func(m *testutils.HandlerDeps) {
				m.UserUseCaseMock.
					EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Return(models.User{
						Id:          1,
						FirstName:   "test",
						LastName:    "test",
						Email:       "test",
						PhoneNumber: "test",
					},
						[]commonutils.ErrorResponse{})
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			deps := testutils.NewHandlerDeps(t)

			// set expectations
			tc.mockFunc(deps)

			w := testutils.DoRequest(t,
				func(r *gin.Engine) {
					h := NewUserHandler(deps.UserUseCaseMock, logger)
					r.GET("/users/get", h.GetUserHandler)
				},
				http.MethodGet, "/users/get", tc.body)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("expected %d, got %d", tc.expectedStatusCode, w.Code)
			}
		})
	}
}

func TestUpdateUserHandler(t *testing.T) {
	logger := testutils.GetLogger()

	// Table of test cases
	testCases := []struct {
		name               string
		body               any
		mockFunc           func(m *testutils.HandlerDeps)
		expectedStatusCode int
	}{
		{
			name: "success - update user",
			body: models.User{
				Id:          1,
				FirstName:   "test",
				LastName:    "test",
				Email:       "test",
				PhoneNumber: "test",
			},
			mockFunc: func(m *testutils.HandlerDeps) {
				m.UserUseCaseMock.
					EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Return(
						[]commonutils.ErrorResponse{})
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			deps := testutils.NewHandlerDeps(t)

			// set expectations
			tc.mockFunc(deps)

			w := testutils.DoRequest(t,
				func(r *gin.Engine) {
					h := NewUserHandler(deps.UserUseCaseMock, logger)
					r.PUT("/users/update", h.UpdateUserHandler)
				},
				http.MethodPut, "/users/update", tc.body)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("expected %d, got %d", tc.expectedStatusCode, w.Code)
			}
		})
	}
}

func TestDeleteUserHandler(t *testing.T) {
	logger := testutils.GetLogger()

	// Table of test cases
	testCases := []struct {
		name               string
		body               any
		mockFunc           func(m *testutils.HandlerDeps)
		expectedStatusCode int
	}{
		{
			name: "success - delete user",
			body: models.IdRequest{
				ID: 1,
			},
			mockFunc: func(m *testutils.HandlerDeps) {
				m.UserUseCaseMock.
					EXPECT().
					DeleteUser(gomock.Any(), gomock.Any()).
					Return(
						[]commonutils.ErrorResponse{})
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			deps := testutils.NewHandlerDeps(t)

			// set expectations
			tc.mockFunc(deps)

			w := testutils.DoRequest(t,
				func(r *gin.Engine) {
					h := NewUserHandler(deps.UserUseCaseMock, logger)
					r.DELETE("/users/delete", h.DeleteUserHandler)
				},
				http.MethodDelete, "/users/delete", tc.body)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("expected %d, got %d", tc.expectedStatusCode, w.Code)
			}
		})
	}
}
