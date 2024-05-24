package main

import (
	"fmt"
	"os"

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
	stratagem stratagem.Model
}

func initialModel() model {
	m := model{
		stratagem: stratagem.New(),
	}

	m.stratagem.Stratagems = []stratagem.Stratagem{
		// ↑ ↓ ← →
		stratagem.NewStratagem("Reinforce", "↑ ↓ → ← ↑", []string{
			` _--_ +`,
			` |<>|  `,
			`_\  /_ `,
		}),
		stratagem.NewStratagem("Hellbomb", "↓ ↑ ← ↓ ↑ → ↓ ↑", []string{
			"(<```>)",
			` _|_|_ `,
			`__| |__`,
		}),
		stratagem.NewStratagem("Orbital Railcannon Strike", "→ ↑ ↓ ↓ →", []string{
			`   X   `,
			` __X__ `,
			`/  V  \`,
		}),
		stratagem.NewStratagem("Eagle 500kg Bomb", "↑ → ↓ ↓ ↓", []string{
            `|\_v_/|`,
            ` \_V_/ `,
            `   V   `,
        }),
	}
	return m
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

		default:
			var cmd tea.Cmd
			m.stratagem, cmd = m.stratagem.Update(msg)
			return m, cmd
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {

	s := "[ @ STRATAGEMS ]\n\n"

	s += m.stratagem.View()

	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
