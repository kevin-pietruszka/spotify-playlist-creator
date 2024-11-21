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

const spotifyGreen = "#008000"

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
		case "esc", "ctrl+c":
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
	leftWidthWithBorder := leftWidth - 2
	rightWidthWithBorder := rightWidth - 2

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(spotifyGreen))

	oneFourthHeight := m.screenHeight / 4
	descriptionHeight := oneFourthHeight - 2
	descriptionBox := borderStyle.
		Width(rightWidthWithBorder).
		Height(descriptionHeight)

	listHeight := m.screenHeight - oneFourthHeight - 2
	listBox := borderStyle.
		Width(rightWidthWithBorder).
		Height(listHeight)

	searchBox := borderStyle.
		Width(leftWidthWithBorder).
		Height(1)

	resultsHeight := m.screenHeight - 2 - 3
	resultsBox := borderStyle.
		Width(leftWidthWithBorder).
		Height(resultsHeight)

	leftPanel := lipgloss.JoinVertical(lipgloss.Top, resultsBox.Render("Results"), searchBox.Render("Search"))
	rightPanel := lipgloss.JoinVertical(lipgloss.Top, listBox.Render("Added List"), descriptionBox.Render("Description"))

	return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)
}

func Run() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
