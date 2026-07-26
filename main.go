package main

import (
	"flag"
	"log"

	"github.com/Akhil373/typr/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	wordCount := flag.Int("w", 30, "number of words")
	flag.Parse()

	if *wordCount < 10 || *wordCount > 100 {
		log.Fatalf("invalid word count %d: must be between 10 - 100", *wordCount)
	}

	p := tea.NewProgram(
		tui.New(*wordCount),
		tea.WithAltScreen(),
	)
	_, err := p.Run()
	if err != nil {
		log.Fatal("error running program", err)
	}
}
