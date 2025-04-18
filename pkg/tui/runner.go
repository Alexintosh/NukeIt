package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// RunTUI launches the TUI for the app uninstall process
func RunTUI(appName string, dryRun, force, verbose bool) error {
	model := NewModel(appName, dryRun, force, verbose)
	
	p := tea.NewProgram(model, tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		return fmt.Errorf("error running TUI: %w", err)
	}
	
	// Get the final model state
	finalModel := m.(Model)
	
	// If there was an error, return it
	if finalModel.errorMsg != "" {
		return fmt.Errorf("%s", finalModel.errorMsg)
	}
	
	return nil
} 