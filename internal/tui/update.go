package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

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
			w, h := m.width, m.height
			m = New(m.wordCount)
			m.width = w
			m.height = h
			return m, nil
		}

	case tickMsg:
		if m.started && !m.finished {
			m.elapsed = time.Since(m.startTime)
			return m, doTick()
		}
	}
	return m, nil
}
