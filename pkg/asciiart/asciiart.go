package asciiart

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

func NewStyled(lines, coloredLines []string, styles map[rune]lipgloss.Style) []string {
	var renderedLines []string

	for lineIdx, line := range lines {
		lineRunes := splitRunes(line)
		coloredRunes := splitRunes(coloredLines[lineIdx])

		var renderedLine string
		for runeIdx, r := range lineRunes {
			styleKey := coloredRunes[runeIdx]
			style, ok := styles[styleKey]
			if ok {
				renderedLine += style.Render(string(r))
			} else {
				renderedLine += string(r)
			}
		}
		renderedLines = append(renderedLines, renderedLine)
	}

	return renderedLines
}

func splitRunes(s string) []rune {
	result := make([]rune, 0, utf8.RuneCountInString(s))
	for _, r := range s {
		result = append(result, r)
	}
	return result
}

func AddBorder(lines []string, style lipgloss.Style) []string {
	width := lipgloss.Width(lines[0])
	result := make([]string, len(lines)+2)
	result[0] = style.Render(fmt.Sprintf("▗%s▖", strings.Repeat("▄", width)))
	for i, line := range lines {
		result[i+1] = fmt.Sprintf("%s%s%s", style.Render("▐"), line, style.Render("▌"))
	}
	result[len(result)-1] = style.Render(fmt.Sprintf("▝%s▘", strings.Repeat("▀", width)))
	return result
}
