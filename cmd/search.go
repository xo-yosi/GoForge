/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xo-yosi/GoForge/internal/apt"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for packages in APT repositories by keyword",
	Long: `Find available packages in APT repositories using a keyword search.

GoForge's search command allows you to discover packages by name, description, or related terms. 
This helps you explore what's available before installing or requesting more information.

Examples:
  goforge search nano         # Find packages related to 'nano'
  goforge search editor       # List editors available in APT
  goforge search python       # Search for Python-related packages

Tip: Use 'goforge info <package>' to see detailed information about a specific package.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("No package specified. Usage: goforge search <package> or goforge search --help")
			return
		}

		pkg := args[0]
		fmt.Printf("Searching for %s...\n", pkg)
		if err := apt.SearchApt(pkg); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
