package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/applejag/kubectl-stratagem/pkg/asciiart"
	"github.com/applejag/kubectl-stratagem/pkg/combo"
	"github.com/applejag/kubectl-stratagem/pkg/stratagem"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	header     []string
	stratagems []stratagem.Model
}

func initialModel() model {
	var (
		white = lipgloss.Color("#f3eedb")
		//beige     = lipgloss.Color("#d6c086")
		darkBeige = lipgloss.Color("#3c3d28")
		red       = lipgloss.Color("#f26e59")
		darkRed   = lipgloss.Color("#4c2b26")
		darkCyan  = lipgloss.Color("#3b4145")
	)
	return model{
		header: asciiart.AddBorder(
			asciiart.NewStyled([]string{
				`   STRATAGEMS `,
			}, []string{
				`wwwwwwwwwwwwwww`,
			}, map[rune]lipgloss.Style{
				'w': lipgloss.NewStyle().Foreground(white).Background(darkCyan).Bold(true),
			}),
			lipgloss.NewStyle().Foreground(darkCyan),
		),
		stratagems: []stratagem.Model{
			// ↑ ↓ ← →
			stratagem.New("Reinforce", "↑ ↓ → ← ↑", "Reinforce deployment by doubling its replicas", asciiart.AddBorder(
				asciiart.NewStyled([]string{
					` _--_ +`,
					` |<>|  `,
					`_\  /_ `,
				}, []string{
					`wwwwwwW`,
					`wwwwwww`,
					`wwwwwww`,
				}, map[rune]lipgloss.Style{
					'w': lipgloss.NewStyle().Foreground(white).Background(darkBeige),
					'W': lipgloss.NewStyle().Foreground(white).Background(darkBeige).Bold(true),
				}),
				lipgloss.NewStyle().Foreground(darkBeige),
			)),

			stratagem.New("Hellbomb", "↓ ↑ ← ↓ ↑ → ↓ ↑", "Liberate an entire namespace (DANGER!)", asciiart.AddBorder(
				asciiart.NewStyled([]string{
					"(<```>)",
					` _|_|_ `,
					`__| |__`,
				}, []string{
					"WWwwwWW",
					`wwWwWww`,
					`wwWwWww`,
				}, map[rune]lipgloss.Style{
					'w': lipgloss.NewStyle().Foreground(white).Background(darkBeige),
					'W': lipgloss.NewStyle().Foreground(white).Background(darkBeige).Bold(true),
				}),
				lipgloss.NewStyle().Foreground(darkBeige),
			)),

			stratagem.New("Orbital Railcannon Strike", "→ ↑ ↓ ↓ →", "Terminate bug pod with Super Destroyer railcannon round", asciiart.AddBorder(
				asciiart.NewStyled([]string{
					`   x   `,
					` __x__ `,
					`/  V  \`,
				}, []string{
					`rrrwrrr`,
					`rrrwrrr`,
					`rrrwrrr`,
				}, map[rune]lipgloss.Style{
					'r': lipgloss.NewStyle().Foreground(red).Background(darkRed),
					'w': lipgloss.NewStyle().Foreground(white).Background(darkRed).Bold(true),
				}),
				lipgloss.NewStyle().Foreground(darkRed),
			)),

			stratagem.New("Eagle 500kg Bomb", "↑ → ↓ ↓ ↓", "Obliterate any bug pods close to impact", asciiart.AddBorder(
				asciiart.NewStyled([]string{
					`|\_v_/|`,
					` \_V_/ `,
					`   V   `,
				}, []string{
					`rrrwrrr`,
					`rrrwrrr`,
					`rrrrrrr`,
				}, map[rune]lipgloss.Style{
					'r': lipgloss.NewStyle().Foreground(red).Background(darkRed),
					'w': lipgloss.NewStyle().Foreground(white).Background(darkRed).Bold(true),
				}),
				lipgloss.NewStyle().Foreground(darkRed),
			)),
		},
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		case "esc":
			return m, combo.ResetCombo

		default:
			return m.updateStrats(msg)
		}

	case combo.ResetComboMsg:
		return m.updateStrats(msg)
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) updateStrats(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	for i, strat := range m.stratagems {
		var cmd tea.Cmd
		m.stratagems[i], cmd = strat.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var sb strings.Builder

	for _, h := range m.header {
		sb.WriteByte(' ')
		sb.WriteString(h)
		sb.WriteByte('\n')
	}

	for _, strat := range m.stratagems {
		sb.WriteString(strat.View())
		//s += "\n"
	}

	sb.WriteString("\nPress q to quit.\n")

	// Send the UI for rendering
	return sb.String()
}
