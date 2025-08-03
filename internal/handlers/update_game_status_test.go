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

type mockUserGameUpdater struct{}

func (m *mockUserGameUpdater) UpdateGame(ctx context.Context, gameUpd models.UserGameUpdate, token string) error {
	return nil
}

func TestUpdateGameStatusHandler_Success(t *testing.T) {
	logger := logging.NewLogger()

	gameUpdate := models.UserGameUpdate{
		GameTitle:  "Cyberpunk",
		GameStatus: "completed",
	}
	body, _ := json.Marshal(gameUpdate)

	req := httptest.NewRequest(http.MethodPatch, "/api/games/status", bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "token", Value: "mock_token"})
	rr := httptest.NewRecorder()

	handler := UpdateGameStatus(*logger, &mockUserGameUpdater{})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
