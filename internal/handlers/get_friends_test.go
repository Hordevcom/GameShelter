package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/stretchr/testify/assert"
)

type mockFriendsGetter struct{}

func (m *mockFriendsGetter) GetFriendsForUserService(ctx context.Context, token string) ([]string, error) {
	return []string{"alice", "bob"}, nil
}

func TestGetFriendsHandler_Success(t *testing.T) {
	logger := logging.NewLogger()

	req := httptest.NewRequest(http.MethodGet, "/api/friends", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: "mock_token"})
	rr := httptest.NewRecorder()

	handler := GetFriendsHandler(*logger, &mockFriendsGetter{})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "alice")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}
