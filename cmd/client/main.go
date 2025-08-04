package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type screen int

const (
	screenSelectAuth screen = iota
	screenAuth
	screenMainMenu
	screenGameList
	screenAddGame
)

type authMode int

const (
	modeLogin authMode = iota
	modeRegister
)

type model struct {
	screen       screen
	authMode     authMode
	username     string
	password     string
	inputFocus   int
	isProcessing bool
	errMsg       string
	token        string
	games        []string
	newGame      string
}

func initialModel() model {
	return model{
		screen: screenSelectAuth,
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Ошибка запуска:", err)
	}
}

// tea.Msg
type authResultMsg struct {
	success bool
	token   string
	err     error
}

type fetchGamesMsg struct {
	games []string
	err   error
}

type addGameMsg struct {
	success bool
	err     error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		key := msg.String()

		switch m.screen {

		case screenSelectAuth:
			switch key {
			case "1":
				m.authMode = modeLogin
				m.screen = screenAuth
			case "2":
				m.authMode = modeRegister
				m.screen = screenAuth
			case "ctrl+c", "q":
				return m, tea.Quit
			}

		case screenAuth:
			switch key {
			case "tab":
				m.inputFocus = (m.inputFocus + 1) % 2
			case "enter":
				if strings.TrimSpace(m.username) == "" || strings.TrimSpace(m.password) == "" {
					m.errMsg = "Username and password required"
					return m, nil
				}
				m.isProcessing = true
				return m, authCmd(m.username, m.password, m.authMode)
			case "backspace":
				if m.inputFocus == 0 && len(m.username) > 0 {
					m.username = m.username[:len(m.username)-1]
				} else if m.inputFocus == 1 && len(m.password) > 0 {
					m.password = m.password[:len(m.password)-1]
				}
			case "esc":
				m.screen = screenSelectAuth
			default:
				if len(key) == 1 {
					if m.inputFocus == 0 {
						m.username += key
					} else {
						m.password += key
					}
				}
			}

		case screenMainMenu:
			switch key {
			case "1":
				return m, fetchGamesCmd(m.token)
			case "2":
				m.newGame = ""
				m.screen = screenAddGame
			case "esc", "q":
				return m, tea.Quit
			}

		case screenAddGame:
			switch key {
			case "enter":
				if strings.TrimSpace(m.newGame) == "" {
					m.errMsg = "Game title required"
					return m, nil
				}
				return m, addGameCmd(m.token, m.newGame)
			case "backspace":
				if len(m.newGame) > 0 {
					m.newGame = m.newGame[:len(m.newGame)-1]
				}
			case "esc":
				m.screen = screenMainMenu
			default:
				if len(key) == 1 {
					m.newGame += key
				}
			}

		case screenGameList:
			if key == "esc" {
				m.screen = screenMainMenu
			}
		}

	case authResultMsg:
		m.isProcessing = false
		if msg.err != nil || !msg.success {
			m.errMsg = "Auth failed"
			return m, nil
		}
		m.token = msg.token
		m.screen = screenMainMenu

	case fetchGamesMsg:
		if msg.err != nil {
			m.errMsg = "Failed to fetch games"
			return m, nil
		}
		m.games = msg.games
		m.screen = screenGameList

	case addGameMsg:
		if msg.err != nil || !msg.success {
			m.errMsg = "Failed to add game"
			return m, nil
		}
		m.errMsg = ""
		m.screen = screenMainMenu
	}

	return m, nil
}

func (m model) View() string {
	switch m.screen {

	case screenSelectAuth:
		return `
GameShelf TUI

1. Login
2. Register

Press Q to quit.
`

	case screenAuth:
		mode := "Login"
		if m.authMode == modeRegister {
			mode = "Register"
		}
		return fmt.Sprintf(
			"%s\n\nUsername: %s\nPassword: %s\n\n[TAB] switch | [ENTER] submit | [ESC] back\n\n%s",
			mode,
			withCursor(m.inputFocus == 0, m.username),
			withCursor(m.inputFocus == 1, strings.Repeat("*", len(m.password))),
			m.errMsg,
		)

	case screenMainMenu:
		return `
Welcome!

1. View My Games
2. Add New Game

Press Q or ESC to quit.
`

	case screenGameList:
		if len(m.games) == 0 {
			return "No games found.\n\n[ESC] back"
		}
		var sb strings.Builder
		sb.WriteString("Your Games:\n\n")
		for _, g := range m.games {
			sb.WriteString("- " + g + "\n")
		}
		sb.WriteString("\n[ESC] back")
		return sb.String()

	case screenAddGame:
		return fmt.Sprintf(
			"Add New Game\n\nTitle: %s\n\n[ENTER] add | [ESC] back\n\n%s",
			m.newGame,
			m.errMsg,
		)

	default:
		return "Unknown screen"
	}
}

func withCursor(focused bool, text string) string {
	if focused {
		return text + "_"
	}
	return text
}

// ─── COMMANDS ─────────────────────────────────────────────

func authCmd(username, password string, mode authMode) tea.Cmd {
	return func() tea.Msg {
		payload := map[string]string{"username": username, "password": password}
		body, _ := json.Marshal(payload)

		url := "http://localhost:8080/api/user/login"
		if mode == modeRegister {
			url = "http://localhost:8080/api/user/register"
		}

		resp, err := http.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			return authResultMsg{success: false, err: err}
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return authResultMsg{success: false}
		}

		for _, c := range resp.Cookies() {
			if c.Name == "token" {
				return authResultMsg{success: true, token: c.Value}
			}
		}
		return authResultMsg{success: false}
	}
}

func fetchGamesCmd(token string) tea.Cmd {
	return func() tea.Msg {
		req, _ := http.NewRequest("GET", "http://localhost:8080/api/users/"+getUsername(token)+"/games", nil)
		req.AddCookie(&http.Cookie{Name: "token", Value: token})
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fetchGamesMsg{err: err}
		}
		defer resp.Body.Close()

		var data []map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return fetchGamesMsg{err: err}
		}

		var titles []string
		for _, g := range data {
			if t, ok := g["game_title"].(string); ok {
				titles = append(titles, t)
			}
		}
		return fetchGamesMsg{games: titles}
	}
}

func addGameCmd(token, title string) tea.Cmd {
	return func() tea.Msg {
		body, _ := json.Marshal(map[string]string{
			"game_title":  title,
			"game_status": "not_owned",
			"game_store":  "steam",
		})
		req, _ := http.NewRequest("POST", "http://localhost:8080/api/users/games", bytes.NewReader(body))
		req.AddCookie(&http.Cookie{Name: "token", Value: token})
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			return addGameMsg{success: false, err: err}
		}
		return addGameMsg{success: true}
	}
}

func getUsername(token string) string {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return ""
	}
	decoded, _ := decodeJWT(parts[1])
	return decoded["Username"]
}

func decodeJWT(payload string) (map[string]string, error) {
	data, err := decodeBase64(payload)
	if err != nil {
		return nil, err
	}
	var result map[string]string
	json.Unmarshal(data, &result)
	return result, nil
}

func decodeBase64(s string) ([]byte, error) {
	s = strings.ReplaceAll(s, "-", "+")
	s = strings.ReplaceAll(s, "_", "/")
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return io.ReadAll(base64.NewDecoder(base64.StdEncoding, strings.NewReader(s)))
}
