package utils

import (
	"reflect"
	"testing"
)

func TestFilterFiles(t *testing.T) {
	testFiles := []File{
		{Path: "/Applications/TestApp.app", Size: 1000, Type: "directory"},
		{Path: "/Users/test/Library/Application Support/com.test.TestApp", Size: 500, Type: "directory"},
		{Path: "/Users/test/Library/Caches/com.test.TestApp", Size: 200, Type: "directory"},
		{Path: "/Users/test/Library/Preferences/com.test.TestApp.plist", Size: 10, Type: "file"},
		{Path: "/Users/test/Documents/TestApp-backup.txt", Size: 5, Type: "file"},
	}

	tests := []struct {
		name       string
		files      []File
		exclusions []string
		want       []File
	}{
		{
			name:       "No exclusions",
			files:      testFiles,
			exclusions: nil,
			want:       testFiles,
		},
		{
			name:       "Empty exclusions",
			files:      testFiles,
			exclusions: []string{},
			want:       testFiles,
		},
		{
			name:       "Exclude by path fragment",
			files:      testFiles,
			exclusions: []string{"Caches"},
			want: []File{
				{Path: "/Applications/TestApp.app", Size: 1000, Type: "directory"},
				{Path: "/Users/test/Library/Application Support/com.test.TestApp", Size: 500, Type: "directory"},
				{Path: "/Users/test/Library/Preferences/com.test.TestApp.plist", Size: 10, Type: "file"},
				{Path: "/Users/test/Documents/TestApp-backup.txt", Size: 5, Type: "file"},
			},
		},
		{
			name:       "Exclude by file name",
			files:      testFiles,
			exclusions: []string{".plist"},
			want: []File{
				{Path: "/Applications/TestApp.app", Size: 1000, Type: "directory"},
				{Path: "/Users/test/Library/Application Support/com.test.TestApp", Size: 500, Type: "directory"},
				{Path: "/Users/test/Library/Caches/com.test.TestApp", Size: 200, Type: "directory"},
				{Path: "/Users/test/Documents/TestApp-backup.txt", Size: 5, Type: "file"},
			},
		},
		{
			name:       "Case insensitive",
			files:      testFiles,
			exclusions: []string{"testapp.app"},
			want: []File{
				{Path: "/Users/test/Library/Application Support/com.test.TestApp", Size: 500, Type: "directory"},
				{Path: "/Users/test/Library/Caches/com.test.TestApp", Size: 200, Type: "directory"},
				{Path: "/Users/test/Library/Preferences/com.test.TestApp.plist", Size: 10, Type: "file"},
				{Path: "/Users/test/Documents/TestApp-backup.txt", Size: 5, Type: "file"},
			},
		},
		{
			name:       "Multiple exclusions",
			files:      testFiles,
			exclusions: []string{"Caches", "Documents"},
			want: []File{
				{Path: "/Applications/TestApp.app", Size: 1000, Type: "directory"},
				{Path: "/Users/test/Library/Application Support/com.test.TestApp", Size: 500, Type: "directory"},
				{Path: "/Users/test/Library/Preferences/com.test.TestApp.plist", Size: 10, Type: "file"},
			},
		},
		{
			name:       "Pattern matching",
			files:      testFiles,
			exclusions: []string{"*.txt"},
			want: []File{
				{Path: "/Applications/TestApp.app", Size: 1000, Type: "directory"},
				{Path: "/Users/test/Library/Application Support/com.test.TestApp", Size: 500, Type: "directory"},
				{Path: "/Users/test/Library/Caches/com.test.TestApp", Size: 200, Type: "directory"},
				{Path: "/Users/test/Library/Preferences/com.test.TestApp.plist", Size: 10, Type: "file"},
			},
		},
		{
			name:       "Exclude everything",
			files:      testFiles,
			exclusions: []string{"/"},
			want:       []File{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterFiles(tt.files, tt.exclusions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterFiles() = %v, want %v", got, tt.want)
			}
		})
	}
} 