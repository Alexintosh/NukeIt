package utils

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestLogVerbose(t *testing.T) {
	buffer := new(bytes.Buffer)
	
	tests := []struct {
		name       string
		verbose    bool
		message    string
		args       []interface{}
		wantOutput string
	}{
		{
			name:       "Verbose enabled",
			verbose:    true,
			message:    "Test message: %s",
			args:       []interface{}{"arg1"},
			wantOutput: "Test message: arg1\n",
		},
		{
			name:       "Verbose disabled",
			verbose:    false,
			message:    "Should not appear: %s",
			args:       []interface{}{"arg2"},
			wantOutput: "",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer.Reset()
			logger := &Logger{
				Verbose: tt.verbose,
				Writer:  buffer,
			}
			
			logger.LogVerbose(tt.message, tt.args...)
			
			if got := buffer.String(); got != tt.wantOutput {
				t.Errorf("LogVerbose() output = %q, want %q", got, tt.wantOutput)
			}
		})
	}
}

func TestLog(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := &Logger{
		Verbose: false, // Should log regardless of verbose setting
		Writer:  buffer,
	}
	
	message := "Test log message: %s"
	args := []interface{}{"arg1"}
	wantOutput := "Test log message: arg1\n"
	
	logger.Log(message, args...)
	
	if got := buffer.String(); got != wantOutput {
		t.Errorf("Log() output = %q, want %q", got, wantOutput)
	}
}

func TestLogResults(t *testing.T) {
	buffer := new(bytes.Buffer)
	
	tests := []struct {
		name       string
		verbose    bool
		results    []DeleteResult
		wantOutput []string // Substrings that should be in output
		unwantOutput []string // Substrings that should not be in output
	}{
		{
			name:    "All successful - Verbose",
			verbose: true,
			results: []DeleteResult{
				{Path: "/path/to/file1", Success: true, Error: nil},
				{Path: "/path/to/file2", Success: true, Error: nil},
			},
			wantOutput: []string{
				"✅ Deleted: /path/to/file1",
				"✅ Deleted: /path/to/file2",
				"Summary: 2 files deleted, 0 failed",
			},
			unwantOutput: []string{
				"Failed to delete",
			},
		},
		{
			name:    "All successful - Not verbose",
			verbose: false,
			results: []DeleteResult{
				{Path: "/path/to/file1", Success: true, Error: nil},
				{Path: "/path/to/file2", Success: true, Error: nil},
			},
			wantOutput: []string{
				"Summary: 2 files deleted, 0 failed",
			},
			unwantOutput: []string{
				"✅ Deleted: /path/to/file1",
				"✅ Deleted: /path/to/file2",
			},
		},
		{
			name:    "Mixed results",
			verbose: true,
			results: []DeleteResult{
				{Path: "/path/to/file1", Success: true, Error: nil},
				{Path: "/path/to/file2", Success: false, Error: errors.New("permission denied")},
			},
			wantOutput: []string{
				"✅ Deleted: /path/to/file1",
				"❌ Failed to delete: /path/to/file2 - permission denied",
				"Summary: 1 files deleted, 1 failed",
			},
			unwantOutput: []string{},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer.Reset()
			logger := &Logger{
				Verbose: tt.verbose,
				Writer:  buffer,
			}
			
			logger.LogResults(tt.results)
			output := buffer.String()
			
			// Check for required substrings
			for _, want := range tt.wantOutput {
				if !strings.Contains(output, want) {
					t.Errorf("LogResults() output should contain %q but got %q", want, output)
				}
			}
			
			// Check for unwanted substrings
			for _, unwant := range tt.unwantOutput {
				if strings.Contains(output, unwant) {
					t.Errorf("LogResults() output should not contain %q but got %q", unwant, output)
				}
			}
		})
	}
} 