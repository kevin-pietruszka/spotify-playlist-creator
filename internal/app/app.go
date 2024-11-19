package app

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	screenHeight int
	screenWidth  int
}

type tickMsg time.Time

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {

	if m.screenHeight < 15 || m.screenWidth < 50 {
		fullscreen := lipgloss.NewStyle().
			Width(m.screenWidth).
			Height(m.screenHeight)
		return fullscreen.Render("Expand your screen")
	}

	rightWidth := m.screenWidth / 3
	leftWidth := m.screenWidth - rightWidth

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#008000"))

	rightPanel := borderStyle.
		Width(rightWidth - 2).
		Height(m.screenHeight - 2)

	leftPanel := borderStyle.
		Width(leftWidth - 2).
		Height(m.screenHeight - 2)

	return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel.Render("Left"), rightPanel.Render("Right"))
}

func Run() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
