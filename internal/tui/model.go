package tui

import (
	"strings"
	"time"

	"github.com/Akhil373/typr/internal/content"
	tea "github.com/charmbracelet/bubbletea"
)

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

func New() model {
	target := strings.Join(content.LoadWords(30), " ")
	return model{
		target: target,
		typed:  "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
