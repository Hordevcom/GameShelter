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

// mockAdder реализует интерфейс FriedsAdder
type mockAdder struct {
	CalledWithToken  string
	CalledWithFriend string
	Err              error
}

func (m *mockAdder) AddFriendForUserService(ctx context.Context, token, friend string) error {
	m.CalledWithToken = token
	m.CalledWithFriend = friend
	return m.Err
}

func TestAddFriendHandler_Success(t *testing.T) {
	friend := models.Friend{Username: "testfriend"}
	body, _ := json.Marshal(friend)

	req := httptest.NewRequest(http.MethodPost, "/api/friends/add", bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "token", Value: "valid_token"})

	rr := httptest.NewRecorder()

	logger := logging.NewLogger()

	mock := &mockAdder{}
	handler := AddFriendHandler(*logger, mock)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"message":"Friend was added!"}`, rr.Body.String())
	assert.Equal(t, "valid_token", mock.CalledWithToken)
	assert.Equal(t, "testfriend", mock.CalledWithFriend)
}
