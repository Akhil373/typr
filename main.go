package main

import (
	"log"

	"github.com/Akhil373/typr/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.New())
	_, err := p.Run()
	if err != nil {
		log.Fatal("error running program", err)
	}
}
