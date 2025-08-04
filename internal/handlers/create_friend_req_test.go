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

type mockFriendReqAdder struct {
	CalledWithPayload models.FriendRequestJSON
	CalledWithToken   string
	Err               error
}

func (m *mockFriendReqAdder) CreateFriendReqService(ctx context.Context, payload models.FriendRequestJSON, token string) error {
	m.CalledWithPayload = payload
	m.CalledWithToken = token
	return m.Err
}

func TestCreateFriendReqHandler_Success(t *testing.T) {
	logger := logging.NewLogger()

	payload := models.FriendRequestJSON{
		Sender:        "user1",
		Receiver:      "user2",
		RequestStatus: "pending",
	}

	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/friendrequest", bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "token", Value: "some_valid_token"})

	rr := httptest.NewRecorder()

	mockService := &mockFriendReqAdder{}
	handler := CreateFriendReqHandler(*logger, mockService)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"message": "Request has created!"}`, rr.Body.String())

	assert.Equal(t, payload, mockService.CalledWithPayload)
	assert.Equal(t, "some_valid_token", mockService.CalledWithToken)
}
