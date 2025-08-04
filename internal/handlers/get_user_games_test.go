package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type mockUserGameGetter struct{}

func (m *mockUserGameGetter) GetUserGames(ctx context.Context, username string) ([]models.UserGames, error) {
	return []models.UserGames{
		{GameTitle: "Cyberpunk", GameStatus: "completed", GameStore: "steam"},
	}, nil
}

func TestGetUserGamesHandler_Success(t *testing.T) {
	logger := logging.NewLogger()

	req := httptest.NewRequest(http.MethodGet, "/api/users/testuser/games", nil)
	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("username", "testuser")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := GetUserGames(*logger, &mockUserGameGetter{})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Cyberpunk")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}
