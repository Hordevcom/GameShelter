package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/stretchr/testify/assert"
)

type mockFriendReqGetter struct{}

func (m *mockFriendReqGetter) GetListOfFriendReqService(ctx context.Context, token string) ([]models.FriendRequestDB, error) {
	return []models.FriendRequestDB{
		{Sender: "alice", Receiver: "bob", RequestStatus: "pending"},
	}, nil
}

func TestGetListOfFriendReqHandler_Success(t *testing.T) {
	logger := logging.NewLogger()

	req := httptest.NewRequest(http.MethodGet, "/api/friendrequests", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: "mock_token"})
	rr := httptest.NewRecorder()

	handler := GetListOfFriendReqHandler(*logger, &mockFriendReqGetter{})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Contains(t, rr.Body.String(), "alice")
}
