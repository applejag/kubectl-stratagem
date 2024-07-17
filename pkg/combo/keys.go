package combo

import "github.com/charmbracelet/bubbles/key"

var DefaultKeyMap = KeyMap{
	StrategemUp: key.NewBinding(
		key.WithKeys("up", "k"),
	),
	StrategemDown: key.NewBinding(
		key.WithKeys("down", "j"),
	),
	StrategemLeft: key.NewBinding(
		key.WithKeys("left", "h"),
	),
	StrategemRight: key.NewBinding(
		key.WithKeys("right", "l"),
	),
}

type KeyMap struct {
	// Keybindings used to input the strategem
	StrategemUp    key.Binding
	StrategemDown  key.Binding
	StrategemLeft  key.Binding
	StrategemRight key.Binding
}
