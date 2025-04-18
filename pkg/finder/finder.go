package finder

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Paths to scan for app-related files
var macOSLibraryPaths = []string{
	"Library/Application Support/",
	"Library/Preferences/",
	"Library/Caches/",
	"Library/Logs/",
	"Library/Containers/",
	"Library/Saved Application State/",
}

type AppFinder struct {
	verbose   bool
	bundleID  string
	appName   string
	foundFiles []string
}

// NewAppFinder creates a new AppFinder instance
func NewAppFinder(verbose bool) *AppFinder {
	return &AppFinder{
		verbose:    verbose,
		foundFiles: make([]string, 0),
	}
}

// FindAllAssociatedFiles finds the app bundle and all associated files
func (f *AppFinder) FindAllAssociatedFiles(appName string) ([]string, error) {
	f.appName = appName
	
	// Search for app bundle in standard locations
	if err := f.findAppBundle(); err != nil {
		return nil, err
	}
	
	// Search for associated files in user's Library
	if err := f.findAssociatedFiles(); err != nil {
		return nil, err
	}
	
	return f.foundFiles, nil
}

// findAppBundle searches for the app bundle in standard locations
func (f *AppFinder) findAppBundle() error {
	// Standard locations for macOS applications
	appLocations := []string{
		"/Applications/",
		filepath.Join(os.Getenv("HOME"), "Applications/"),
	}
	
	appFound := false
	
	for _, location := range appLocations {
		// Try with standard .app extension
		appPath := filepath.Join(location, f.appName+".app")
		if f.verbose {
			fmt.Printf("Checking for app bundle at: %s\n", appPath)
		}
		
		if _, err := os.Stat(appPath); err == nil {
			// App found, add to found files
			f.foundFiles = append(f.foundFiles, appPath)
			appFound = true
			
			// Try to extract bundle ID
			bundleID, err := f.extractBundleID(appPath)
			if err != nil {
				if f.verbose {
					fmt.Printf("Warning: Could not extract bundle ID: %v\n", err)
				}
			} else {
				f.bundleID = bundleID
				if f.verbose {
					fmt.Printf("Found bundle ID: %s\n", f.bundleID)
				}
			}
			
			break
		}
		
		// If not found, try without .app extension (for non-standard app directories)
		appPath = filepath.Join(location, f.appName)
		if f.verbose {
			fmt.Printf("Checking for app directory at: %s\n", appPath)
		}
		
		if info, err := os.Stat(appPath); err == nil && info.IsDir() {
			// App directory found, add to found files
			f.foundFiles = append(f.foundFiles, appPath)
			appFound = true
			
			// For non-standard app directories, we may not be able to extract bundle ID
			// but we'll try using a common naming pattern
			f.bundleID = fmt.Sprintf("com.%s.%s", strings.ToLower(f.appName), strings.ToLower(f.appName))
			if f.verbose {
				fmt.Printf("Non-standard app directory found. Using assumed bundle ID: %s\n", f.bundleID)
			}
			
			break
		}
	}
	
	if !appFound && f.verbose {
		fmt.Printf("App bundle not found for %s\n", f.appName)
	}
	
	return nil
}

// extractBundleID extracts the bundle ID from the app's Info.plist
func (f *AppFinder) extractBundleID(appPath string) (string, error) {
	// Use our plist parser to extract the bundle ID
	bundleID, err := f.ParseBundleID(appPath)
	if err != nil {
		// Fallback to a generic bundle ID format if parsing fails
		return fmt.Sprintf("com.example.%s", strings.ToLower(f.appName)), fmt.Errorf("failed to parse bundle ID: %w", err)
	}
	return bundleID, nil
}

// findAssociatedFiles searches for app-related files in standard macOS directories
func (f *AppFinder) findAssociatedFiles() error {
	homeDir := os.Getenv("HOME")
	
	for _, libPath := range macOSLibraryPaths {
		fullPath := filepath.Join(homeDir, libPath)
		if f.verbose {
			fmt.Printf("Scanning directory: %s\n", fullPath)
		}
		
		// Walk through the directory and find matches
		if err := filepath.WalkDir(fullPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				// Skip directories we can't access
				if f.verbose {
					fmt.Printf("Warning: Could not access %s: %v\n", path, err)
				}
				return filepath.SkipDir
			}
			
			// Skip the root directory
			if path == fullPath {
				return nil
			}
			
			// Check if the file/directory matches our app
			if f.isRelatedToApp(path) {
				f.foundFiles = append(f.foundFiles, path)
				if f.verbose {
					fmt.Printf("Found related file: %s\n", path)
				}
				
				// If it's a directory, no need to scan its contents individually
				if d.IsDir() {
					return filepath.SkipDir
				}
			}
			
			return nil
		}); err != nil {
			if f.verbose {
				fmt.Printf("Warning: Error scanning %s: %v\n", fullPath, err)
			}
		}
	}
	
	return nil
}

// isRelatedToApp checks if a file or directory is related to the app
func (f *AppFinder) isRelatedToApp(path string) bool {
	baseName := filepath.Base(path)
	
	// Check for bundle ID match
	if f.bundleID != "" && strings.Contains(baseName, f.bundleID) {
		return true
	}
	
	// Check for app name match (case-insensitive)
	appNameLower := strings.ToLower(f.appName)
	baseNameLower := strings.ToLower(baseName)
	
	return strings.Contains(baseNameLower, appNameLower)
} 