package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
)

var sentences []string = []string{
	// Easy
	"the cat sat on the mat",
	"I like to eat apples",
	"sunny days are nice",
	"let's go to the café",

	// Medium
	"she asked me if I do this everyday",
	"hello, how are you doing",
	"the quick brown fox jumps over the lazy dog",
	"practice makes perfect every single day",
}

type tickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type model struct {
	target    string
	typed     string
	startTime time.Time
	started   bool
	finished  bool
	elapsed   time.Duration
}

func initialModel() model {
	target := sentences[rand.Intn(len(sentences))]
	return model{
		target: target,
		typed:  "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyBackspace:
			if len(m.typed) > 0 {
				m.typed = m.typed[:len(m.typed)-1]
			}

		case tea.KeyRunes, tea.KeySpace:
			if m.finished {
				return m, nil
			}
			var cmd tea.Cmd
			if !m.started {
				m.started = true
				m.startTime = time.Now()
				cmd = doTick()
			}
			m.typed += msg.String()

			if len(m.typed) >= len(m.target) {
				m.finished = true
				m.elapsed = time.Since(m.startTime)
				return m, nil
			}
			return m, cmd

		case tea.KeyCtrlR:
			return initialModel(), nil
		}

	case tickMsg:
		if m.started && !m.finished {
			m.elapsed = time.Since(m.startTime)
			return m, doTick()
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.finished {
		charCount := utf8.RuneCountInString(m.target)
		timeMinutes := m.elapsed.Seconds() / 60.0
		wpm := (float64(charCount) / 5.0) / timeMinutes

		targetWords := strings.Split(m.target, " ")
		typedWords := strings.Split(m.typed, " ")
		correctWords := 0
		minLen := min(len(targetWords), len(typedWords))

		for i := range minLen {
			if targetWords[i] == typedWords[i] {
				correctWords++
			}
		}
		accuracy := float64(correctWords) / float64(len(targetWords)) * 100
		return fmt.Sprintf(
			"\n--- Results ---\nElapsed Time: %.1fs\nAccuracy: %.2f%%\nWPM: %.2f\n\nPress ctrl-r to restart or ctrl- to quit.\n",
			m.elapsed.Seconds(), accuracy, wpm,
		)
	}

	var b strings.Builder
	b.WriteString("Type the text below:\n\n")

	targetRunes := []rune(m.target)
	typedRunes := []rune(m.typed)

	for i, ch := range targetRunes {
		if i < len(typedRunes) {
			if typedRunes[i] == ch {
				fmt.Fprintf(&b, "\033[32m%c\033[0m", typedRunes[i])
			} else {
				fmt.Fprintf(&b, "\033[31m%c\033[0m", ch)
			}
		} else if i == len(typedRunes) {
			fmt.Fprintf(&b, "\033[4m%c\033[0m", ch)
		} else {
			b.WriteString(string(ch))
		}
	}
	fmt.Fprintf(&b, "\n\nTime: %.1fs\n", m.elapsed.Seconds())
	b.WriteString("(Press ctrl-r to restart or ctrl-c to quit.")

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}
}
