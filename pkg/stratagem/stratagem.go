package stratagem

import (
	"slices"
	"strings"

	"github.com/applejag/kubectl-stratagem/pkg/combo"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var DefaultStyle = Style{
	Name: StyleName{
		Idle:    lipgloss.NewStyle().Foreground(lipgloss.Color("#f3eedb")).Bold(true),
		Wrong:   lipgloss.NewStyle().Foreground(lipgloss.Color("#827f74")),
		Correct: lipgloss.NewStyle().Foreground(lipgloss.Color("#59f258")).Bold(true),
	},
	Desc: StyleDesc{
		Idle:    lipgloss.NewStyle().Foreground(lipgloss.Color("#827f74")).Italic(true),
		Wrong:   lipgloss.NewStyle().Foreground(lipgloss.Color("#827f74")).Italic(true),
		Correct: lipgloss.NewStyle().Foreground(lipgloss.Color("#f3eedb")).Italic(true),
	},
}

func New(name, comboStr, desc string, art []string) Model {
	return Model{
		Name:  name,
		Desc:  desc,
		Art:   art,
		Combo: combo.New(combo.NewCombo(comboStr)),
		Style: DefaultStyle,
	}
}

type Model struct {
	Name  string
	Desc  string
	Art   []string
	Combo combo.Model
	Style Style
}

type Style struct {
	Name StyleName
	Desc StyleDesc
}

type StyleName struct {
	Idle    lipgloss.Style
	Wrong   lipgloss.Style
	Correct lipgloss.Style
}

type StyleDesc struct {
	Idle    lipgloss.Style
	Wrong   lipgloss.Style
	Correct lipgloss.Style
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		var cmd tea.Cmd
		m.Combo, cmd = m.Combo.Update(msg)
		return m, cmd

	case combo.ResetComboMsg:
		var cmd tea.Cmd
		m.Combo, cmd = m.Combo.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	lines := slices.Clone(m.Art)
	for i := range lines {
		lines[i] = " " + lines[i]
	}
	middle := len(lines) / 2

	switch m.Combo.State {
	case combo.StateCorrect:
		lines[middle-1] += "  " + m.Style.Name.Correct.Render(strings.ToUpper(m.Name))
		lines[middle-0] += "  " + m.Style.Desc.Correct.Render(m.Desc)
	case combo.StateWrong:
		lines[middle-1] += "  " + m.Style.Name.Wrong.Render(strings.ToUpper(m.Name))
		lines[middle-0] += "  " + m.Style.Desc.Wrong.Render(m.Desc)
	default:
		lines[middle-1] += "  " + m.Style.Name.Idle.Render(strings.ToUpper(m.Name))
		lines[middle-0] += "  " + m.Style.Desc.Idle.Render(m.Desc)
	}

	lines[middle+1] += "  " + m.Combo.View()

	return strings.Join(lines, "\n") + "\n"
}
