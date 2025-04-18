package finder

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// SimplePlist represents the basic structure of a plist file
type SimplePlist struct {
	XMLName xml.Name  `xml:"plist"`
	Dict    PlistDict `xml:"dict"`
}

// PlistDict represents a dictionary in a plist file
type PlistDict struct {
	Items []xml.Name `xml:",any"`
}

// PlistItem is an interface for plist values
type PlistItem interface{}

// PlistString represents a string in a plist
type PlistString struct {
	Value string `xml:",chardata"`
}

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
	
	// For testing simplicity, let's do a simple search for CFBundleIdentifier in the XML
	// This is more reliable than trying to parse the XML which can have different formats
	contentStr := string(content)
	
	// Look for the pattern: <key>CFBundleIdentifier</key>\s*<string>VALUE</string>
	bundleIDStart := strings.Index(contentStr, "<key>CFBundleIdentifier</key>")
	if bundleIDStart == -1 {
		return "", fmt.Errorf("bundle ID not found in Info.plist")
	}
	
	// Find the opening <string> tag after the key
	stringTagStart := strings.Index(contentStr[bundleIDStart:], "<string>")
	if stringTagStart == -1 {
		return "", fmt.Errorf("bundle ID value not found in Info.plist")
	}
	
	// Adjust the position to account for the substring operation
	stringTagStart += bundleIDStart + len("<string>")
	
	// Find the closing </string> tag
	stringTagEnd := strings.Index(contentStr[stringTagStart:], "</string>")
	if stringTagEnd == -1 {
		return "", fmt.Errorf("bundle ID value not properly formatted in Info.plist")
	}
	
	// Extract the bundle ID
	bundleID := contentStr[stringTagStart : stringTagStart+stringTagEnd]
	
	if bundleID == "" {
		return "", fmt.Errorf("empty bundle ID in Info.plist")
	}
	
	return bundleID, nil
} 