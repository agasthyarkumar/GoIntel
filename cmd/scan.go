package cmd

import (
	"fmt"
	"log"
	"runtime"
	"sync"

	"gointel/internal/analytics"
	"gointel/internal/models"
	"gointel/internal/scanner"
	"gointel/internal/worker"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:     "scan [path]",
	Aliases: []string{"s"},
	Short:   "Scan files and detect duplicates",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		root := args[0]

		jobs := make(chan string, 100)
		results := make(chan models.FileInfo, 100)

		var wg sync.WaitGroup

		duplicates := models.DuplicateMap{}
		var allFiles []models.FileInfo

		var mu sync.Mutex

		numWorkers := runtime.NumCPU()

		fmt.Println("🚀 Starting workers:", numWorkers)

		// Workers
		for i := 1; i <= numWorkers; i++ {

			wg.Add(1)

			go func(id int) {
				defer wg.Done()

				worker.StartWorker(
					id,
					jobs,
					results,
				)

			}(i)
		}

		// Scanner
		go func() {

			err := scanner.ScanDirectory(root, jobs)
			if err != nil {
				log.Println(err)
			}

		}()

		// Close results
		go func() {
			wg.Wait()
			close(results)
		}()

		// Process results
		for result := range results {

			fmt.Println("================================")
			fmt.Println("File:", result.Path)
			fmt.Println("Size:", result.Size, "bytes")
			fmt.Println("SHA256:", result.Checksum)

			mu.Lock()

			allFiles = append(allFiles, result)

			duplicates[result.Checksum] = append(
				duplicates[result.Checksum],
				result.Path,
			)

			mu.Unlock()
		}

		// Duplicate Report
		fmt.Println("\n==============================")
		fmt.Println("🔥 DUPLICATE FILES")
		fmt.Println("==============================")

		found := false

		for hash, files := range duplicates {

			if len(files) > 1 {

				found = true

				fmt.Println("\nHash:", hash)

				for _, file := range files {
					fmt.Println(" -", file)
				}
			}
		}

		if !found {
			fmt.Println("✅ No duplicate files found")
		}

		// Analytics
		fmt.Println("\n==============================")
		fmt.Println("📊 STORAGE ANALYTICS")
		fmt.Println("==============================")

		var totalSize int64

		for _, file := range allFiles {
			totalSize += file.Size
		}

		fmt.Println("Total Files:", len(allFiles))
		fmt.Println("Total Size:", totalSize, "bytes")

		largest := analytics.GetLargestFiles(allFiles, 5)

		fmt.Println("\n🔥 Largest Files")

		for _, file := range largest {

			fmt.Println("--------------------------------")
			fmt.Println("File:", file.Path)
			fmt.Println("Size:", file.Size, "bytes")
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}