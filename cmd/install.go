/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a package from APT or GitHub releases",
	Long: `Install packages from multiple sources automatically.

GoForge detects the best source for each package:
- APT packages: Install directly from system repositories
- GitHub releases: Download and install from GitHub releases

The package manager automatically handles:
- Source detection and selection
- Download and installation
- Dependency resolution (for APT packages)
- Binary placement in ~/.goforge/bin/
- Installation tracking and logging

Examples:
  goforge install vim          # Install vim from APT
  goforge install code         # Install VS Code from APT  
  goforge install gh           # Install GitHub CLI from releases
  goforge install docker       # Install Docker from APT
  goforge install bat          # Install bat from GitHub releases
  goforge install node        # Install Node.js from APT

Use 'goforge list' to see all installed packages.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
