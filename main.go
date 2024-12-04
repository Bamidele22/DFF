package main

import (
	"fmt"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
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
	// Implementation 
	var wg sync.WaitGroup
	filesMap := make(map[string][]string)

	// Traverse the directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err!= nil {
			return err
		}

		// Calculate MD5 hash for file content
		hasher := md5.New()
		file, err := os.Open(path)
		if err!= nil {
			return err
		}
		defer file.Close()

		if _, err := io.Copy(hasher, file); err!= nil {
			return err
		}

		contentHash := hex.EncodeToString(hasher.Sum(nil))
		filesMap[contentHash] = append(filesMap[contentHash], path)

		wg.Add(1)
		go func() {
			defer wg.Done()
			checkDuplicates(contentHash, filesMap)
		}()

		return nil
	})

	if err!= nil {
		fmt.Println(err)
		return
	}

	wg.Wait()
}

func checkDuplicates(contentHash string, filesMap map[string][]string) {
	if len(filesMap[contentHash]) > 1 {
		for _, filePath := range filesMap[contentHash] {
			fmt.Printf("Duplicate found: %s\n", filePath)
		}
	}

}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
