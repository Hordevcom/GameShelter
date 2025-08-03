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

type mockFriendReqRemover struct{}

func (m *mockFriendReqRemover) DeleteFriendReqService(ctx context.Context, payload models.FriendRequestJSON, token string) error {
	return nil
}

func TestDeleteFriendReqHandler_Success(t *testing.T) {
	logger := logging.NewLogger()

	body, _ := json.Marshal(models.FriendRequestJSON{
		Sender:        "alice",
		Receiver:      "bob",
		RequestStatus: "pending",
	})
	req := httptest.NewRequest(http.MethodDelete, "/api/friendrequest", bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "token", Value: "mock_token"})
	rr := httptest.NewRecorder()

	handler := DeleteFriendReqHandler(*logger, &mockFriendReqRemover{})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Request was deleted")
}
