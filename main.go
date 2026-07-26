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
	p := tea.NewProgram(
		tui.New(*wordCount),
		tea.WithAltScreen(),
	)
	_, err := p.Run()
	if err != nil {
		log.Fatal("error running program", err)
	}
}
