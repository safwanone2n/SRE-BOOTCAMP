package testutils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	usermock "rest-api/internal/mocks/user"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/remiges-tech/logharbour/logharbour"
)

type HandlerDeps struct {
	Ctrl            *gomock.Controller
	UserUseCaseMock *usermock.MockUserUseCasesI
}

func NewHandlerDeps(t *testing.T) *HandlerDeps {
	t.Helper()
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	return &HandlerDeps{
		Ctrl:            ctrl,
		UserUseCaseMock: usermock.NewMockUserUseCasesI(ctrl),
	}
}

// generic helper to perform a request and return status code
func DoRequest(t *testing.T, register func(*gin.Engine), method, path string, body any) *httptest.ResponseRecorder {
	t.Helper()

	r := gin.New()
	register(r)

	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}
	}

	req := httptest.NewRequest(method, path, &buf)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	return w
}

func GetLogger() *logharbour.Logger {
	return logharbour.NewLogger(logharbour.NewLoggerContext(logharbour.Info), "REST-API", io.Discard)
}
