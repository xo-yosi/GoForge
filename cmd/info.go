/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xo-yosi/GoForge/internal/apt"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show detailed information about an APT package",
	Long: `Display detailed information for a package from APT repositories.

GoForge's info command retrieves and displays key details about a package,
including its version, maintainer, dependencies, and description.
This helps you verify package details before installation.

Examples:
  goforge info nano         # Show information about the 'nano' package
  goforge info git          # Get details for the 'git' package
  goforge info python3      # View info for the 'python3' package

Tip: Use 'goforge search <keyword>' to find available packages by keyword.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("No package specified. Usage: goforge info <package> or goforge info --help")
			return
		}

		pkg := args[0]
		if err := apt.InfoApt(pkg); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}