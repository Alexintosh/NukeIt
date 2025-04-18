package cleaner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CriticalPaths lists critical system paths that should never be touched
var criticalPaths = []string{
	"/System",
	"/bin",
	"/sbin",
	"/usr/bin",
	"/usr/sbin",
	"/usr/local/bin",
	"/usr/local/sbin",
	"/etc",
	"/var",
}

// SafeDirs are directories that are safe to remove app-related files from
var safeDirs = []string{
	"Library/Application Support",
	"Library/Preferences",
	"Library/Caches",
	"Library/Logs",
	"Library/Containers",
	"Library/Saved Application State",
}

// UnsafeDirs are directories that we should never remove files from
var unsafeDirs = []string{
	"Documents",
	"Downloads",
	"Desktop",
	"Pictures",
	"Music",
	"Movies",
}

type AppCleaner struct {
	verbose bool
}

// NewAppCleaner creates a new AppCleaner instance
func NewAppCleaner(verbose bool) *AppCleaner {
	return &AppCleaner{
		verbose: verbose,
	}
}

// DeleteFiles safely deletes the list of provided files
func (c *AppCleaner) DeleteFiles(files []string) (int, error) {
	deleted := 0
	
	for _, file := range files {
		if !c.isSafeToDelete(file) {
			if c.verbose {
				fmt.Printf("Skipping potentially unsafe path: %s\n", file)
			}
			continue
		}
		
		if c.verbose {
			fmt.Printf("Deleting: %s\n", file)
		}
		
		if err := os.RemoveAll(file); err != nil {
			fmt.Printf("Error deleting %s: %v\n", file, err)
		} else {
			deleted++
		}
	}
	
	return deleted, nil
}

// isSafeToDelete checks if a file or directory is safe to delete
func (c *AppCleaner) isSafeToDelete(path string) bool {
	// Check if the path is a critical system path
	for _, criticalPath := range criticalPaths {
		if strings.HasPrefix(path, criticalPath) {
			return false
		}
	}
	
	// Check if the path is in an unsafe directory
	homeDir := os.Getenv("HOME")
	for _, unsafeDir := range unsafeDirs {
		unsafePath := filepath.Join(homeDir, unsafeDir)
		if strings.HasPrefix(path, unsafePath) {
			return false
		}
	}
	
	// Check if the path is in a safe directory
	inSafeDir := false
	for _, safeDir := range safeDirs {
		safePath := filepath.Join(homeDir, safeDir)
		if strings.HasPrefix(path, safePath) {
			inSafeDir = true
			break
		}
	}
	
	// Applications directories are also safe
	if strings.HasPrefix(path, "/Applications/") || strings.HasPrefix(path, filepath.Join(homeDir, "Applications/")) {
		inSafeDir = true
	}
	
	return inSafeDir
} 