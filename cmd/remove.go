/*
Copyright © 2025 NAME HERE xoyosi19@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Uninstall a package installed by GoForge",
Long: `Remove packages previously installed with GoForge from your system.

GoForge's remove command safely uninstalls packages that were installed via APT or GitHub releases
using GoForge. This helps you keep your environment clean and organized.

The package manager automatically handles:
- Detecting the package source (APT or GitHub)
- Uninstalling via the appropriate method
- Cleaning up binaries from ~/.goforge/bin/
- Updating the installation tracking database

Examples:
  goforge remove nano         # Uninstall nano installed from APT
  goforge remove gh           # Remove GitHub CLI installed from releases
  goforge remove bat          # Remove bat installed from GitHub releases

Tip: Use 'goforge list' to see all packages you can remove.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
