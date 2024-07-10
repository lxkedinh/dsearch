/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/lxkedinh/dsearch/search"
	"github.com/lxkedinh/dsearch/spinner"
	"github.com/spf13/cobra"
)

var startDir, searchQuery string
var fuzzyThreshold float64
var progressIndicator = spinner.New()

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "dsearch path",
	Short:   "dsearch is a Go CLI application to quickly fuzzy search your system.",
	Long:    `dsearch is a Go CLI application to quickly fuzzy search your system.`,
	Aliases: []string{"ds"},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := ValidateFlags()
		if err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		progressIndicator.Start()
		results, err := search.Fuzzy(startDir, searchQuery, fuzzyThreshold)
		progressIndicator.Stop()
		if err != nil {
			fmt.Printf("Unexpected error occurred. Try again.\nError: %v", err)
			return err
		}

		for _, result := range results {
			fmt.Printf("• %s\n", result)
		}

		return nil
	},
	// TODO figure out design of arguments
	Args: ValidateArgs,
}

func ValidateFlags() error {
	err := ValidateStartDir(startDir)
	if err != nil {
		return err
	}

	err = ValidateFuzzyThreshold(fuzzyThreshold)
	if err != nil {
		return err
	}

	return nil
}

func ValidateFuzzyThreshold(threshold float64) error {
	if fuzzyThreshold < 0.0 || fuzzyThreshold > 1.0 {
		return errors.New("invalid threshold value, must be between 0.0 and 1.0")
	}

	return nil
}

func ValidateArgs(cmd *cobra.Command, args []string) error {
	switch len(args) {
	case 1:
		searchQuery = args[0]
	default:
		return errors.New("incorrect arguments, only 1 argument for search query")
	}

	return nil
}

func ValidateStartDir(startDir string) error {
	fileInfo, err := os.Stat(startDir)
	if errors.Is(err, fs.ErrNotExist) || !fileInfo.IsDir() {
		return errors.New("invalid starting directory given")
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dsearch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.Flags().StringVarP(&startDir, "dir", "d", ".", "Starting directory to start fuzzy searching from (defaults to current directory)")
	RootCmd.Flags().Float64VarP(&fuzzyThreshold, "threshold", "t", 0.5, "Threshold to determine strictness of fuzzy matching (allows values from 0.0 to 1.0)")
}
