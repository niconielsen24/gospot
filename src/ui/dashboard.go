package ui

import (
	"fmt"
  tea "github.com/charmbracelet/bubbletea"
)


type DashboardModel struct {
	credentials Credentials
}

func NewDashboardModel(credentials Credentials) DashboardModel {
	return DashboardModel{credentials: credentials}
}

func (m DashboardModel) Init() tea.Cmd {
	return nil
}

func (m DashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m DashboardModel) View() string {
	return fmt.Sprintf("Welcome, %s!\n\nPress q to quit.", m.credentials.Username)
}

