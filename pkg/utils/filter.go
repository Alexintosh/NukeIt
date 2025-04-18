package utils

import (
	"path/filepath"
	"strings"
)

// File represents a file to be processed by the application
type File struct {
	Path string
	Size int64
	Type string // file, directory, etc.
}

// FilterFiles removes files that match any of the exclusion patterns
func FilterFiles(files []File, exclusions []string) []File {
	if len(exclusions) == 0 {
		return files
	}
	
	filtered := make([]File, 0, len(files))
	
	for _, file := range files {
		excluded := false
		
		for _, exclusion := range exclusions {
			// Convert both to lower case for case-insensitive comparison
			// This is important on macOS which has a case-insensitive filesystem
			if strings.Contains(strings.ToLower(file.Path), strings.ToLower(exclusion)) {
				excluded = true
				break
			}
			
			// Also check if the exclusion matches as a path pattern
			if matched, _ := filepath.Match(exclusion, filepath.Base(file.Path)); matched {
				excluded = true
				break
			}
		}
		
		if !excluded {
			filtered = append(filtered, file)
		}
	}
	
	return filtered
} 