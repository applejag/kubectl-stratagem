package main

import (
	"fmt"
	"os"

	"github.com/applejag/kubectl-stratagem/pkg/combo"
	"github.com/applejag/kubectl-stratagem/pkg/stratagem"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	stratagems []stratagem.Model
}

func initialModel() model {
	return model{
		stratagems: []stratagem.Model{
			// ↑ ↓ ← →
			stratagem.New("Reinforce", "↑ ↓ → ← ↑", []string{
				` _--_ +`,
				` |<>|  `,
				`_\  /_ `,
			}),
			stratagem.New("Hellbomb", "↓ ↑ ← ↓ ↑ → ↓ ↑", []string{
				"(<```>)",
				` _|_|_ `,
				`__| |__`,
			}),
			stratagem.New("Orbital Railcannon Strike", "→ ↑ ↓ ↓ →", []string{
				`   x   `,
				` __x__ `,
				`/  V  \`,
			}),
			stratagem.New("Eagle 500kg Bomb", "↑ → ↓ ↓ ↓", []string{
				`|\_v_/|`,
				` \_V_/ `,
				`   V   `,
			}),
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

	s := "[ @ STRATAGEMS ]\n\n"

	for _, strat := range m.stratagems {
		s += strat.View()
		s += "\n"
	}

	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
