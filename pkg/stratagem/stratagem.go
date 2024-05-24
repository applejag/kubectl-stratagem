package stratagem

import (
	"fmt"
	"strings"

	"github.com/applejag/kubectl-stratagem/pkg/combo"
	tea "github.com/charmbracelet/bubbletea"
)

func New(name, comboStr string, art []string) Model {
	return Model{
		Name:  name,
		Art:   art,
		Combo: combo.New(combo.NewCombo(comboStr), combo.CharsetNerdFontBold),
	}
}

type Model struct {
	Name  string
	Art   []string
	Combo combo.Model
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
	var sb strings.Builder


	fmt.Fprintf(&sb, ""+
		" %s   %s\n"+
		" %s  \n"+
		" %s   %s\n",
		m.Art[0], strings.ToUpper(m.Name),
		m.Art[1],
		m.Art[2], m.Combo.View(),
	)

	return sb.String()
}
