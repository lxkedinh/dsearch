/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var startDir, searchPath string

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "dsearch path",
	Short:   "dsearch is a Go CLI application to quickly fuzzy search your system.",
	Long:    `dsearch is a Go CLI application to quickly fuzzy search your system.`,
	Aliases: []string{"ds"},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ValidateStartDir(startDir)
		if err != nil {
			return err
		}

		fmt.Println(startDir)

		filepath.WalkDir(startDir, func(path string, d fs.DirEntry, err error) error {
			if fs.DirEntry.IsDir(d) {
				fmt.Println(d)
			}
			return nil
		})

		return nil
	},
	// TODO figure out design of arguments
	Args: ValidateArgs,
}

func ValidateArgs(cmd *cobra.Command, args []string) error {
	switch len(args) {
	case 1:
		err := ValidateSearchPath(args[0])
		if err != nil {
			return err
		}
	default:
		return errors.New("incorrect arguments, only 1 argument for desired search path")
	}

	return nil
}

func ValidateSearchPath(p string) error {
	path, err := filepath.Abs(p)
	if err != nil {
		return errors.New("invalid file path given for fuzzy searching")
	}
	searchPath = path
	return nil
}

func ValidateStartDir(startDir string) error {
	fileInfo, err := os.Stat(startDir)
	if errors.Is(err, fs.ErrNotExist) || !fileInfo.IsDir() {
		return errors.New("invalid starting directory given")
	}

	// change file path to use "/" instead of "\" separator to work regardless of OS
	startDir = filepath.ToSlash(startDir)

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
}
