package finder

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// SimplePlist represents the basic structure of a plist file
type SimplePlist struct {
	Dict PlistDict `xml:"dict"`
}

// PlistDict represents a dictionary in a plist file
type PlistDict struct {
	Keys   []string    `xml:"key"`
	Values []PlistItem `xml:",any"`
}

// PlistItem is an interface for plist values
type PlistItem interface{}

// ParseBundleID extracts the bundle ID from an app's Info.plist file
func (f *AppFinder) ParseBundleID(appPath string) (string, error) {
	infoPlistPath := filepath.Join(appPath, "Contents", "Info.plist")
	
	// Check if Info.plist exists
	if _, err := os.Stat(infoPlistPath); err != nil {
		return "", fmt.Errorf("Info.plist not found: %w", err)
	}
	
	// Open the file
	file, err := os.Open(infoPlistPath)
	if err != nil {
		return "", fmt.Errorf("failed to open Info.plist: %w", err)
	}
	defer file.Close()
	
	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read Info.plist: %w", err)
	}
	
	// Parse XML
	var plist SimplePlist
	if err := xml.Unmarshal(content, &plist); err != nil {
		return "", fmt.Errorf("failed to parse Info.plist: %w", err)
	}
	
	// Find the bundle ID
	for i, key := range plist.Dict.Keys {
		if key == "CFBundleIdentifier" && i < len(plist.Dict.Values) {
			if bundleID, ok := plist.Dict.Values[i+1].(string); ok {
				return bundleID, nil
			}
		}
	}
	
	return "", fmt.Errorf("bundle ID not found in Info.plist")
} 