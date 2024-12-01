package app

import (
	"log"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	screenHeight int
	screenWidth  int
	leftWidth    int
	rightWidth   int
	textInput    textinput.Model
}

type tickMsg time.Time

const spotifyGreen = "#008000"

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Type to look for artists, playlist, etc."
	ti.Focus()
	ti.CharLimit = 128
	return model{
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width

		m.rightWidth = m.screenWidth / 3
		m.leftWidth = m.screenWidth - m.rightWidth

		m.textInput.Width = m.leftWidth - 2 - 2 - 1
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(message)

	return m, cmd
}

func (m model) View() string {

	if m.screenHeight < 15 || m.screenWidth < 50 {
		fullscreen := lipgloss.NewStyle().
			Width(m.screenWidth).
			Height(m.screenHeight)
		return fullscreen.Render("Expand your screen")
	}

	leftWidthWithBorder := m.leftWidth - 2
	rightWidthWithBorder := m.rightWidth - 2

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(spotifyGreen))

	oneThirdHeight := m.screenHeight / 3
	descriptionHeight := oneThirdHeight - 2
	descriptionBox := borderStyle.
		Width(rightWidthWithBorder).
		Height(descriptionHeight)

	listHeight := m.screenHeight - oneThirdHeight - 2
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

	leftPanel := lipgloss.JoinVertical(lipgloss.Top, resultsBox.Render("Results"), searchBox.Render(m.textInput.View()))
	rightPanel := lipgloss.JoinVertical(lipgloss.Top, listBox.Render("Added List"), descriptionBox.Render("Description"))

	return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)
}

func Run() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
