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

type mockGameStorage struct {
	Exist bool
	Err   error
}

func (m *mockGameStorage) IsGameAlreadyExist(ctx context.Context, gamename string) (error, bool) {
	return m.Err, m.Exist
}

func (m *mockGameStorage) AddNewGame(ctx context.Context, game models.Game) error {
	return nil
}

type mockUserGameAdder struct {
	Exist bool
	Err   error
}

func (m *mockUserGameAdder) CheckGameInUserLib(ctx context.Context, gametitle string, token string) (error, bool) {
	return m.Err, m.Exist
}

func (m *mockUserGameAdder) AddGameToUser(ctx context.Context, usergame models.UserGameJSON, token string) error {
	return m.Err
}

func TestAddGameToUser_Success(t *testing.T) {
	logger := logging.NewLogger()

	game := models.UserGameJSON{
		GameTitle:  "Cyberpunk 2077",
		GameStatus: "completed",
		GameStore:  "steam",
	}
	body, _ := json.Marshal(game)

	req := httptest.NewRequest(http.MethodPost, "/api/users/games", bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "token", Value: "valid_token"})

	rr := httptest.NewRecorder()

	handler := AddGameToUser(
		*logger,
		&mockGameStorage{Exist: true},
		&mockUserGameAdder{Exist: false},
	)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"message":"Game added to user library"}`, rr.Body.String())
}
