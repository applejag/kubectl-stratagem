package combo

import (
	"fmt"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var DefaultStyle = Style{
	Idle:           lipgloss.NewStyle().Foreground(lipgloss.Color("#f3eedb")),
	Input:          lipgloss.NewStyle().Foreground(lipgloss.Color("#827f74")).Bold(true),
	Correct:        lipgloss.NewStyle().Foreground(lipgloss.Color("#59f258")),
	Wrong:          lipgloss.NewStyle().Foreground(lipgloss.Color("#f26e59")),
	WrongRemaining: lipgloss.NewStyle().Foreground(lipgloss.Color("#827f74")),
}

type Style struct {
	Idle           lipgloss.Style
	Input          lipgloss.Style
	Correct        lipgloss.Style
	Wrong          lipgloss.Style
	WrongRemaining lipgloss.Style
}

type ResetComboMsg struct{}

func ResetCombo() tea.Msg {
	return ResetComboMsg{}
}

type ComboWrongMsg struct {
	Model Model
}

type ComboCorrectMsg struct {
	Model Model
}

type State byte

const (
	StateIdle State = iota
	StateWrong
	StateCorrect
)

type Key rune

const (
	KeyUp    Key = '🡅'
	KeyDown  Key = '🡇'
	KeyLeft  Key = '🡄'
	KeyRight Key = '🡆'
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

func NewCombo(s string) Combo {
	combo := make(Combo, 0, len(s)-strings.Count(s, " "))
	for _, key := range s {
		switch key {
		case rune(KeyUp),
			rune(KeyDown),
			rune(KeyLeft),
			rune(KeyRight):
			combo = append(combo, Key(key))
		case ' ':
			// ignore
		default:
			panic(fmt.Errorf("invalid stratagem key: %q", key))
		}
	}
	return combo
}

func New(combo Combo) Model {
	return Model{
		Combo: combo,
		Style: DefaultStyle,
	}
}

type Model struct {
	Combo Combo
	Input int
	State State
	Style Style
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			return m.updateInput(KeyUp)
		case "down":
			return m.updateInput(KeyDown)
		case "left":
			return m.updateInput(KeyLeft)
		case "right":
			return m.updateInput(KeyRight)

		default:
			return m, nil
		}

	case ResetComboMsg:
		m.Input = 0
		m.State = StateIdle
		return m, nil
	}

	return m, nil
}

func (m Model) updateInput(key Key) (Model, tea.Cmd) {
	if m.State != StateIdle {
		return m, nil
	}

	if m.Combo[m.Input] != key {
		m.State = StateWrong
		m.Input++
		return m, m.comboWrong()
	}

	m.Input++

	if m.Input >= len(m.Combo) {
		m.State = StateCorrect
		return m, m.comboCorrect()
	}

	return m, nil
}

func (m Model) View() string {

	switch m.State {
	case StateCorrect:
		return m.Style.Correct.Render(m.Combo.String() )
	case StateWrong:
		matching, nonMatching := m.splitCombo()
		if len(matching) == 0 {
			return m.Style.WrongRemaining.Render(nonMatching.String() )
		}
		if len(nonMatching) == 0 {
			return m.Style.Wrong.Render(matching.String() )
		}
		return m.Style.Wrong.Render(matching.String()+" ") +
			m.Style.WrongRemaining.Render(nonMatching.String())
	default:
		matching, nonMatching := m.splitCombo()
		if len(matching) == 0 {
			return m.Style.Idle.Render(nonMatching.String() )
		}
		if len(nonMatching) == 0 {
			return m.Style.Idle.Render(matching.String() )
		}
		return m.Style.Input.Render(matching.String()+" ") +
			m.Style.Idle.Render(nonMatching.String())
	}
}

func (m Model) splitCombo() (matching, nonMatching Combo) {
	if m.Input <= 0 {
		return nil, m.Combo
	}

	if m.Input >= len(m.Combo) {
		return m.Combo, nil
	}

	return m.Combo[:m.Input], m.Combo[m.Input:]
}

func (m Model) comboCorrect() tea.Cmd {
	return func() tea.Msg {
		return ComboCorrectMsg{m}
	}
}

func (m Model) comboWrong() tea.Cmd {
	return func() tea.Msg {
		return ComboWrongMsg{m}
	}
}
