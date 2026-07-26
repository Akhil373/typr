package tui

import (
	"fmt"
	"strings"

	"github.com/Akhil373/typr/internal/stats"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}
	var b strings.Builder

	boxStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#CDD6F4")).
		Padding(1, 2)

	if m.finished {
		res := stats.Compute(m.target, m.typed, m.elapsed)
		headerStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true)

		labelStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#89B4FA"))

		mutedStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6C7086")).
			Italic(true)

		b.WriteString("\n" + headerStyle.Render("--- Results ---") + "\n")

		fmt.Fprintf(
			&b,
			"%s %.1fs\n%s %.2f%%\n%s %.2f\n\n%s\n",
			labelStyle.Render("Elapsed:"), res.Elapsed.Seconds(),
			labelStyle.Render("Accuracy:"), res.Accuracy,
			labelStyle.Render("WPM:"), res.WPM,
			mutedStyle.Render("Press ctrl-r to restart or ctrl-c to quit."),
		)

		styledContent := boxStyle.Render(b.String())

		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			styledContent,
			lipgloss.WithWhitespaceChars(" "),
			lipgloss.WithWhitespaceForeground(lipgloss.Color("#1E1E2E")),
		)
	}

	b.WriteString("Type the text below:\n\n")

	words := strings.Fields(m.target)
	typedRunes := []rune(m.typed)
	pos := 0

	for wi, word := range words {
		if wi > 0 && wi%5 == 0 {
			b.WriteByte('\n')
		}
		var (
			correctStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6EE065"))
			wrongStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#E04E4C"))
			cursorStyle  = lipgloss.NewStyle().Underline(true)
		)
		for _, ch := range word {
			if pos < len(typedRunes) {
				if typedRunes[pos] == ch {
					b.WriteString(correctStyle.Render(string(typedRunes[pos])))
				} else {
					b.WriteString(wrongStyle.Render(string(ch)))
				}
			} else if pos == len(typedRunes) {
				b.WriteString(cursorStyle.Render(string(ch)))
			} else {
				b.WriteRune(ch)
			}
			pos++
		}

		if wi < len(words)-1 {
			if pos < len(typedRunes) {
				if typedRunes[pos] == ' ' {
					b.WriteString(correctStyle.Render(" "))
				} else {
					b.WriteString(wrongStyle.Render(" "))
				}
			} else if pos == len(typedRunes) {
				b.WriteString(cursorStyle.Render(" "))
			} else {
				b.WriteByte(' ')
			}
			pos++
		}
	}
	fmt.Fprintf(&b, "\n\nTime: %.1fs\n", m.elapsed.Seconds())
	b.WriteString("(Press ctrl-r to restart or ctrl-c to quit.)")

	styledContent := boxStyle.Render(b.String())

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		styledContent,
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("#1E1E2E")),
	)
}
