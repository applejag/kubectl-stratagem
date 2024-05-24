package stratagem

import (
	"fmt"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Key rune

const (
	KeyUp    Key = '↑'
	KeyDown  Key = '↓'
	KeyLeft  Key = '←'
	KeyRight Key = '→'
)

type Combo []Key

func (c Combo) Equal(other Combo) bool {
	return slices.Equal(c, other)
}

func (c Combo) HasPrefix(prefix Combo) bool {
	return len(c) >= len(prefix) && c[0:len(prefix)].Equal(prefix)
}

func (c Combo) CutPrefix(prefix Combo) (Combo, bool) {
	if !c.HasPrefix(prefix) {
		return c, false
	}
	return c[len(prefix):], true
}

func (c Combo) String() string {
	if len(c) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.Grow(len(c)*2 - 1)
	for i, key := range c {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteRune(rune(key))
	}
	return sb.String()
}

type Stratagem struct {
	Name  string
	Combo Combo
	Art   []string
}

func NewStratagem(name, combo string, art []string) Stratagem {
	strat := Stratagem{
		Name:  name,
		Combo: make(Combo, 0, len(combo)),
		Art:   art,
	}
	for _, key := range combo {
		switch key {
		case rune(KeyUp),
			rune(KeyDown),
			rune(KeyLeft),
			rune(KeyRight):
			strat.Combo = append(strat.Combo, Key(key))
		case ' ':
			// ignore
		default:
			panic(fmt.Errorf("invalid stratagem key: %q", key))
		}
	}
	return strat
}

type Model struct {
	Input      Combo
	Stratagems []Stratagem
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.addInput(KeyUp)
		case "down":
			m.addInput(KeyDown)
		case "left":
			m.addInput(KeyLeft)
		case "right":
			m.addInput(KeyRight)

		case "esc":
			m.Input = nil

		default:
			return m, nil
		}
	}

	return m, nil
}

func (m *Model) addInput(key Key) {
	m.Input = append(m.Input, key)

	if !m.validInput(m.Input) {
		m.Input = nil
	}

	for _, strat := range m.Stratagems {
		if strat.Combo.Equal(m.Input) {
			// woohoo!
			m.Input = nil
		}
	}
}

func (m Model) validInput(combo Combo) bool {
	for _, strat := range m.Stratagems {
		if strat.Combo.HasPrefix(combo) {
			return true
		}
	}
	return false
}

func (m Model) View() string {
	var sb strings.Builder

	for _, strat := range m.Stratagems {
		nonMatching, isMatching := strat.Combo.CutPrefix(m.Input)
		matching := m.Input
		if !isMatching {
			matching = nil
		}

		fmt.Fprintf(&sb, ""+
			"%s | %s\n"+
			"%s |\n"+
			"%s | %s%q\n",
			strat.Art[0], strings.ToUpper(strat.Name),
			strat.Art[1],
			strat.Art[2], matching, nonMatching,
		)

		sb.WriteByte('\n')
	}

	return sb.String()
}
