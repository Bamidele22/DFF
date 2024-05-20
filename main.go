package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "duplicate-finder",
	Short: "Find and remove duplicate files in a directory.",
	Long:  `Duplicate Finder is a command-line utility that scans directories for duplicate files and provides an option to delete them.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a directory to scan.")
			return
		}

		directory := args[0]
		findDuplicates(directory)
	},
}

func findDuplicates(dir string) {
	// Implementation goes here

}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
