package tui

import (
	"fmt"
	"strings"

	"github.com/Akhil373/typr/internal/stats"
)

func (m model) View() string {
	if m.finished {
		res := stats.Compute(m.target, m.typed, m.elapsed)
		return fmt.Sprintf(
			"\n--- Results ---\nElapsed Time: %.1fs\nAccuracy: %.2f%%\nWPM: %.2f\n\nPress ctrl-r to restart or ctrl-c to quit.\n",
			res.Elapsed.Seconds(), res.Accuracy, res.WPM,
		)
	}

	var b strings.Builder
	b.WriteString("Type the text below:\n\n")

	words := strings.Fields(m.target)
	typedRunes := []rune(m.typed)
	pos := 0

	for wi, word := range words {
		if wi > 0 && wi%5 == 0 {
			b.WriteByte('\n')
		}
		for _, ch := range word {
			if pos < len(typedRunes) {
				if typedRunes[pos] == ch {
					fmt.Fprintf(&b, "\033[32m%c\033[0m", typedRunes[pos])
				} else {
					fmt.Fprintf(&b, "\033[31m%c\033[0m", ch)
				}
			} else if pos == len(typedRunes) {
				fmt.Fprintf(&b, "\033[4m%c\033[0m", ch)
			} else {
				b.WriteRune(ch)
			}
			pos++
		}
		if wi < len(words)-1 {
			if pos < len(typedRunes) {
				if typedRunes[pos] == ' ' {
					fmt.Fprintf(&b, "\033[32m \033[0m")
				} else {
					fmt.Fprintf(&b, "\033[31m \033[0m")
				}
			} else if pos == len(typedRunes) {
				fmt.Fprintf(&b, "\033[4m \033[0m")
			} else {
				b.WriteByte(' ')
			}
			pos++
		}
	}
	fmt.Fprintf(&b, "\n\nTime: %.1fs\n", m.elapsed.Seconds())
	b.WriteString("(Press ctrl-r to restart or ctrl-c to quit.)")

	return b.String()
}
