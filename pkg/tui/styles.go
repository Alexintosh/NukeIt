package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	danger    = lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"}

	// Styles
	appStyle = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(highlight)

	titleStyle = lipgloss.NewStyle().
		Foreground(highlight).
		Bold(true).
		Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
		Foreground(special)

	errorStyle = lipgloss.NewStyle().
		Foreground(danger)

	spinnerStyle = lipgloss.NewStyle().
		Foreground(highlight)

	fileListStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(subtle).
		Padding(0, 1)

	selectedItemStyle = lipgloss.NewStyle().
		Foreground(highlight).
		Bold(true)

	checkboxChecked = lipgloss.NewStyle().
		SetString("[✓] ").
		Foreground(special)

	checkboxUnchecked = lipgloss.NewStyle().
		SetString("[ ] ").
		Foreground(subtle)

	progressBarStyle = lipgloss.NewStyle().
		Foreground(highlight).
		Bold(true)

	progressBarFilled = lipgloss.NewStyle().
		Background(highlight).
		Foreground(lipgloss.Color("#ffffff"))

	progressBarEmpty = lipgloss.NewStyle().
		Background(subtle)

	summaryStyle = lipgloss.NewStyle().
		Padding(1, 0).
		Foreground(special).
		Bold(true)
) 