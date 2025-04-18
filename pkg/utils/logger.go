package utils

import (
	"fmt"
	"io"
	"os"
)

// DeleteResult represents the result of a file deletion operation
type DeleteResult struct {
	Path    string
	Success bool
	Error   error
}

// Logger handles logging of operations
type Logger struct {
	Verbose bool
	Writer  io.Writer
}

// NewLogger creates a new Logger instance
func NewLogger(verbose bool) *Logger {
	return &Logger{
		Verbose: verbose,
		Writer:  os.Stdout,
	}
}

// LogVerbose logs a message only if verbose mode is enabled
func (l *Logger) LogVerbose(format string, args ...interface{}) {
	if l.Verbose {
		fmt.Fprintf(l.Writer, format+"\n", args...)
	}
}

// Log logs a message regardless of verbose setting
func (l *Logger) Log(format string, args ...interface{}) {
	fmt.Fprintf(l.Writer, format+"\n", args...)
}

// LogResults logs the results of file deletion operations
func (l *Logger) LogResults(results []DeleteResult) {
	success := 0
	failed := 0
	
	for _, result := range results {
		if result.Success {
			success++
			if l.Verbose {
				l.Log("✅ Deleted: %s", result.Path)
			}
		} else {
			failed++
			l.Log("❌ Failed to delete: %s - %v", result.Path, result.Error)
		}
	}
	
	l.Log("\nSummary: %d files deleted, %d failed", success, failed)
} 