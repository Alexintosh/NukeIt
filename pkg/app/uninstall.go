package app

import (
	"fmt"
	"strings"

	"github.com/alexintosh/gocleaner/pkg/cleaner"
	"github.com/alexintosh/gocleaner/pkg/finder"
	"github.com/spf13/cobra"
)

var (
	dryRun  bool
	force   bool
	verbose bool
)

func init() {
	uninstallCmd := &cobra.Command{
		Use:   "uninstall <AppName>",
		Short: "Uninstall an application and its associated files",
		Long: `Uninstall an application by removing the main app bundle and associated files
like caches, preferences, logs, etc. from various macOS system paths.`,
		Args: cobra.ExactArgs(1),
		RunE: runUninstall,
	}

	uninstallCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would be deleted, but don't delete")
	uninstallCmd.Flags().BoolVar(&force, "force", false, "Skip confirmation prompt and delete files immediately")
	uninstallCmd.Flags().BoolVar(&verbose, "verbose", false, "Show detailed scanning and deletion info")

	rootCmd.AddCommand(uninstallCmd)
}

func runUninstall(cmd *cobra.Command, args []string) error {
	appName := args[0]
	
	// Remove .app suffix if provided - we'll handle both cases in the finder
	appName = strings.TrimSuffix(appName, ".app")

	// Find app bundle and associated files
	appFinder := finder.NewAppFinder(verbose)
	foundFiles, err := appFinder.FindAllAssociatedFiles(appName)
	if err != nil {
		return fmt.Errorf("error finding files: %w", err)
	}

	if len(foundFiles) == 0 {
		fmt.Printf("No files found for %s\n", appName)
		return nil
	}

	// Print found files
	fmt.Printf("Found %d files associated with %s:\n", len(foundFiles), appName)
	for _, file := range foundFiles {
		fmt.Printf("- %s\n", file)
	}

	// If dry run, exit here
	if dryRun {
		fmt.Println("\nThis was a dry run. No files were deleted.")
		return nil
	}

	// Confirm deletion unless force flag is set
	if !force {
		fmt.Print("\nAre you sure you want to delete these files? (y/N): ")
		var confirm string
		fmt.Scanln(&confirm)
		if strings.ToLower(confirm) != "y" {
			fmt.Println("Operation cancelled.")
			return nil
		}
	}

	// Delete files
	appCleaner := cleaner.NewAppCleaner(verbose)
	deleted, err := appCleaner.DeleteFiles(foundFiles)
	if err != nil {
		return fmt.Errorf("error deleting files: %w", err)
	}

	fmt.Printf("\nSuccessfully deleted %d files.\n", deleted)
	return nil
} 