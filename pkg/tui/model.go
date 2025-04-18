package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/alexintosh/gocleaner/pkg/cleaner"
	"github.com/alexintosh/gocleaner/pkg/finder"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Model states
const (
	stateScanning   = "scanning"
	stateSelectFiles = "select_files"
	stateDeleting   = "deleting"
	stateDone       = "done"
)

// FileItem represents a file in the list
type FileItem struct {
	path     string
	selected bool
}

func (i FileItem) Title() string {
	if i.selected {
		return checkboxChecked.String() + i.path
	}
	return checkboxUnchecked.String() + i.path
}

func (i FileItem) Description() string {
	return ""
}

func (i FileItem) FilterValue() string {
	return i.path
}

// Model represents the TUI state
type Model struct {
	appName      string
	dryRun       bool
	force        bool
	verbose      bool
	state        string
	spinner      spinner.Model
	fileList     list.Model
	progress     progress.Model
	files        []string
	selectedFiles []string
	errorMsg     string
	statusMsg    string
	appFinder    *finder.AppFinder
	appCleaner   *cleaner.AppCleaner
	width        int
	height       int
}

// NewModel creates a new TUI model
func NewModel(appName string, dryRun, force, verbose bool) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	p := progress.New(progress.WithDefaultGradient())
	p.ShowPercentage = true

	fileList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	fileList.Title = "Files to be removed:"
	fileList.SetFilteringEnabled(false)
	fileList.SetShowStatusBar(false)
	fileList.SetShowHelp(true)
	fileList.Styles.Title = titleStyle

	return Model{
		appName:    appName,
		dryRun:     dryRun,
		force:      force,
		verbose:    verbose,
		state:      stateScanning,
		spinner:    s,
		progress:   p,
		fileList:   fileList,
		appFinder:  finder.NewAppFinder(verbose),
		appCleaner: cleaner.NewAppCleaner(verbose),
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.scanFiles,
	)
}

// scanFiles searches for files associated with the app
func (m Model) scanFiles() tea.Msg {
	files, err := m.appFinder.FindAllAssociatedFiles(m.appName)
	if err != nil {
		return errMsg{err}
	}

	return filesFoundMsg{files}
}

// Update handles UI state changes
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.fileList.SetSize(msg.Width-4, msg.Height-10)
		m.progress.Width = msg.Width - 10
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case " ":
			if m.state == stateSelectFiles {
				index := m.fileList.Index()
				if index >= 0 && index < len(m.files) {
					fileItems := []list.Item{}
					for i, item := range m.fileList.Items() {
						fileItem := item.(FileItem)
						if i == index {
							fileItem.selected = !fileItem.selected
						}
						fileItems = append(fileItems, fileItem)
					}
					m.fileList.SetItems(fileItems)
				}
			}
			return m, nil

		case "enter":
			if m.state == stateSelectFiles {
				// Get selected files
				selectedFiles := []string{}
				for _, item := range m.fileList.Items() {
					fileItem := item.(FileItem)
					if fileItem.selected {
						selectedFiles = append(selectedFiles, fileItem.path)
					}
				}
				m.selectedFiles = selectedFiles

				if len(selectedFiles) == 0 {
					return m, tea.Quit
				}

				// If dry run, just exit
				if m.dryRun {
					m.state = stateDone
					m.statusMsg = fmt.Sprintf("Dry run complete. %d files would be deleted.", len(selectedFiles))
					return m, tea.Quit
				}

				// If force or if we confirmed, start deleting
				if m.force {
					m.state = stateDeleting
					return m, m.startDeleting
				}

				m.state = stateDeleting
				return m, m.startDeleting
			}
			return m, nil

		case "a":
			if m.state == stateSelectFiles {
				fileItems := []list.Item{}
				for _, item := range m.fileList.Items() {
					fileItem := item.(FileItem)
					fileItem.selected = true
					fileItems = append(fileItems, fileItem)
				}
				m.fileList.SetItems(fileItems)
			}
			return m, nil

		case "n":
			if m.state == stateSelectFiles {
				fileItems := []list.Item{}
				for _, item := range m.fileList.Items() {
					fileItem := item.(FileItem)
					fileItem.selected = false
					fileItems = append(fileItems, fileItem)
				}
				m.fileList.SetItems(fileItems)
			}
			return m, nil
		}

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case filesFoundMsg:
		m.files = msg.files
		m.state = stateSelectFiles

		fileItems := []list.Item{}
		for _, file := range m.files {
			fileItems = append(fileItems, FileItem{
				path:     file,
				selected: true,
			})
		}
		m.fileList.SetItems(fileItems)

		// If no files found
		if len(m.files) == 0 {
			m.state = stateDone
			m.statusMsg = fmt.Sprintf("No files found for %s", m.appName)
			return m, tea.Quit
		}

		// If force flag is set, skip selection and go straight to deletion
		if m.force {
			m.selectedFiles = m.files
			m.state = stateDeleting
			return m, m.startDeleting
		}

		return m, nil

	case errMsg:
		m.state = stateDone
		m.errorMsg = msg.err.Error()
		return m, tea.Quit

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case progressUpdateMsg:
		if msg.done {
			m.state = stateDone
			m.statusMsg = fmt.Sprintf("Successfully deleted %d files.", msg.count)
			return m, tea.Quit
		}
		cmd := m.progress.SetPercent(float64(msg.current) / float64(msg.total))
		return m, cmd
	}

	// Update list when in select files state
	if m.state == stateSelectFiles {
		var listCmd tea.Cmd
		m.fileList, listCmd = m.fileList.Update(msg)
		return m, listCmd
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	var s strings.Builder

	switch m.state {
	case stateScanning:
		s.WriteString(m.spinner.View())
		s.WriteString(" Scanning for files associated with ")
		s.WriteString(titleStyle.Render(m.appName))
		s.WriteString("...\n")

	case stateSelectFiles:
		s.WriteString(titleStyle.Render(fmt.Sprintf("Found %d files for %s\n\n", len(m.files), m.appName)))
		s.WriteString(fileListStyle.Render(m.fileList.View()))
		s.WriteString("\nUse arrow keys to navigate, space to toggle selection, a to select all, n to select none\n")
		s.WriteString("Press Enter to confirm or Ctrl+C to quit\n")

	case stateDeleting:
		s.WriteString(titleStyle.Render("Deleting files...\n\n"))
		s.WriteString(m.progress.View() + "\n\n")

	case stateDone:
		if m.errorMsg != "" {
			s.WriteString(errorStyle.Render("Error: " + m.errorMsg + "\n"))
		} else {
			s.WriteString(statusMessageStyle.Render(m.statusMsg + "\n"))
		}
	}

	return appStyle.Render(s.String())
}

// startDeleting begins the file deletion process
func (m Model) startDeleting() tea.Msg {
	totalFiles := len(m.selectedFiles)
	deletedCount := 0

	// Send initial progress update
	time.Sleep(100 * time.Millisecond)
	tea.Tick(time.Millisecond*10, func(t time.Time) tea.Msg {
		return progressUpdateMsg{current: 0, total: totalFiles, count: 0, done: false}
	})

	// Delete files
	for i, file := range m.selectedFiles {
		if m.appCleaner.IsSafeToDelete(file) {
			err := m.appCleaner.DeleteSingleFile(file)
			if err == nil {
				deletedCount++
			}
		}

		// Update progress
		time.Sleep(100 * time.Millisecond)
		tea.Tick(time.Millisecond*10, func(t time.Time) tea.Msg {
			return progressUpdateMsg{current: i + 1, total: totalFiles, count: deletedCount, done: false}
		})
	}

	// Final update
	return progressUpdateMsg{current: totalFiles, total: totalFiles, count: deletedCount, done: true}
}

// Messages
type filesFoundMsg struct {
	files []string
}

type errMsg struct {
	err error
}

func (e errMsg) Error() string { return e.err.Error() }

type progressUpdateMsg struct {
	current int
	total   int
	count   int
	done    bool
} 