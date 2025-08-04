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

type mockFriendRemover struct{}

func (m *mockFriendRemover) DeleteFriendForUserService(ctx context.Context, token, friend string) error {
	return nil
}

func TestDeleteFriendHandler_Success(t *testing.T) {
	logger := logging.NewLogger()

	body, _ := json.Marshal(models.Friend{Username: "bob"})
	req := httptest.NewRequest(http.MethodDelete, "/api/friends", bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "token", Value: "mock_token"})
	rr := httptest.NewRecorder()

	handler := DeleteFriendHandler(*logger, &mockFriendRemover{})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Friend was deleted")
}
