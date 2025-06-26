/*
Copyright © 2025 NAME HERE xoyosi19@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xo-yosi/GoForge/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "GoForge",
	Short: "GoForge - Universal Package Manager for Linux",
	Long: `GoForge is a universal package manager that unifies APT packages and GitHub releases.

Install, manage, and track packages from multiple sources with a single command.
GoForge automatically detects package sources and handles installations seamlessly.

Examples:
  goforge install vim          # Install from APT
  goforge install gh           # Install GitHub CLI from releases  
  goforge list                 # Show installed packages
  goforge search nodejs        # Search across all sources
  goforge remove vim           # Remove any installed package

GoForge creates ~/.goforge/ to store configuration and track installations.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
    cobra.OnInitialize(initConfig)
}

func initConfig() {
    if err := config.Init(); err != nil {
        fmt.Printf("Error initializing config: %v\n", err)
        os.Exit(1)
    }
}
