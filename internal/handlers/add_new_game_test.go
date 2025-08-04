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

type mockGameStorageAdder struct {
	GameExists bool
	Err        error
	AddedGame  *models.Game
}

func (m *mockGameStorageAdder) IsGameAlreadyExist(ctx context.Context, gamename string) (error, bool) {
	return m.Err, m.GameExists
}

func (m *mockGameStorageAdder) AddNewGame(ctx context.Context, game models.Game) error {
	m.AddedGame = &game
	return m.Err
}

func TestAddNewGame_Success(t *testing.T) {
	logger := logging.NewLogger()

	game := models.Game{
		Title: "Hollow Knight",
		Genre: "Metroidvania",
	}
	body, _ := json.Marshal(game)

	req := httptest.NewRequest(http.MethodPost, "/api/games", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	mockStorage := &mockGameStorageAdder{
		GameExists: false,
	}

	handler := AddNewGame(*logger, mockStorage)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"message": "Game added to server library"}`, rr.Body.String())
	assert.NotNil(t, mockStorage.AddedGame)
	assert.Equal(t, game.Title, mockStorage.AddedGame.Title)
	assert.Equal(t, game.Genre, mockStorage.AddedGame.Genre)
}
