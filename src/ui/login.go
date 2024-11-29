package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Credentials struct {
  Username string
  Password string
}

type model struct {
	cursor   int
	username textinput.Model
	password textinput.Model
}

func InitialModel() model {
	un := textinput.New()
	un.Prompt = "| "
	un.Placeholder = "Your username"
	un.CharLimit = 20
  un.Focus()
	pw := textinput.New()
	pw.Prompt = "| "
	pw.Placeholder = "Your password"
	pw.EchoMode = textinput.EchoPassword
	pw.EchoCharacter = '*'
	pw.CharLimit = 20
	return model{
		cursor:   0,
		username: un,
		password: pw,
	}
}

func (m model) Init() tea.Cmd {
	return nil 
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "ctrl+k":
			if m.cursor > 0 {
				m.cursor--
				m.username.Focus()
				m.password.Blur()
			}
		case "down", "ctrl+j":
			if m.cursor < 1 {
				m.cursor++
				m.password.Focus()
				m.username.Blur()
			}
		case "enter", " ":
			if m.username.Value() != "" && m.password.Value() != "" {
        var new_cred = Credentials{
          Username: m.username.Value(),
          Password: m.password.Value(),
        }
        return NewDashboardModel(new_cred), nil
      }
      return m, tea.Quit
		case "esc":
			m.username.Blur()
			m.password.Blur()
			m.username.Reset()
			m.password.Reset()
		}

		if m.cursor == 0 {
			var cmd tea.Cmd
			m.username, cmd = m.username.Update(msg)
			return m, cmd
		} else {
			var cmd tea.Cmd
			m.password, cmd = m.password.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Go spot login\n\n"

	if m.cursor == 0 {
		s += fmt.Sprintf("> %s\n  %s\n", m.username.View(), m.password.View())
	} else {
		s += fmt.Sprintf("  %s\n> %s\n", m.username.View(), m.password.View())
	}

	s += "\n\nPress [q, ctrl+c] to quit\nPress r to reset fields\n"

	return s
}
