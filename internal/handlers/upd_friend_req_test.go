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

type mockFriendReqUpdater struct{}

func (m *mockFriendReqUpdater) UpdateFriendReqService(ctx context.Context, payload models.FriendRequestJSON, token string) error {
	return nil
}

func TestUpdateFriendReqHandler_Success(t *testing.T) {
	logger := logging.NewLogger()

	reqBody := models.FriendRequestJSON{
		Sender:        "alice",
		Receiver:      "bob",
		RequestStatus: "accepted",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPatch, "/api/friendrequest", bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "token", Value: "mock_token"})
	rr := httptest.NewRecorder()

	handler := UpdateFriendReqHandler(*logger, &mockFriendReqUpdater{})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Request was updated")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}
