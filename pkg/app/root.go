package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nuke",
	Short: "A CLI tool to fully uninstall macOS applications",
	Long: `nuke is a command-line tool for macOS that helps users fully uninstall applications 
by removing the main app bundle and associated files like caches, preferences, logs, etc.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
} 