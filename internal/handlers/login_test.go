package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/stretchr/testify/assert"
)

type mockUserAuthLogin struct{}

func (m *mockUserAuthLogin) CheckUserPassword(ctx context.Context, user models.UserAuth) error {
	return nil
}

func (m *mockUserAuthLogin) CheckUserLoginService(ctx context.Context, user models.UserAuth) bool {
	return false
}

func TestLoginHandler_Success(t *testing.T) {
	logger := logging.NewLogger()

	user := models.UserAuth{
		Username: "john_doe",
		Password: "secret123",
	}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler := Login(&mockUserAuthLogin{}, *logger)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	cookies := rr.Result().Cookies()
	foundToken := false
	for _, c := range cookies {
		if c.Name == "token" {
			foundToken = true
			break
		}
	}
	assert.True(t, foundToken, "Expected token cookie to be set")
}
