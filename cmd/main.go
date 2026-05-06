package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	"gointel/internal/analytics"
	"gointel/internal/crypto"
	"gointel/internal/models"
	"gointel/internal/scanner"
	"gointel/internal/worker"
)

func main() {

	// ========================================
	// CHECK ARGUMENTS
	// ========================================

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("gointel <directory>")
		fmt.Println("gointel encrypt <file>")
		fmt.Println("gointel decrypt <file>")
		return
	}

	command := os.Args[1]

	// ========================================
	// ENCRYPTION MODE
	// ========================================

	if command == "encrypt" {

		if len(os.Args) < 3 {
			fmt.Println("Usage:")
			fmt.Println("gointel encrypt <file>")
			return
		}

		file := os.Args[2]

		fmt.Println("🔐 Encrypting:", file)

		err := crypto.EncryptFile(file)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("✅ Encryption complete")

		return
	}

	// ========================================
	// DECRYPTION MODE
	// ========================================

	if command == "decrypt" {

		if len(os.Args) < 3 {
			fmt.Println("Usage:")
			fmt.Println("gointel decrypt <file>")
			return
		}

		file := os.Args[2]

		fmt.Println("🔓 Decrypting:", file)

		err := crypto.DecryptFile(file)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("✅ Decryption complete")

		return
	}

	// ========================================
	// SCAN MODE
	// ========================================

	root := command

	// Channels
	jobs := make(chan string, 100)
	results := make(chan models.FileInfo, 100)

	// WaitGroup
	var wg sync.WaitGroup

	// Duplicate detector
	duplicates := models.DuplicateMap{}

	// Store all processed files
	var allFiles []models.FileInfo

	// Mutex for shared state
	var mu sync.Mutex

	// Worker count
	numWorkers := runtime.NumCPU()

	fmt.Println("🚀 Starting workers:", numWorkers)

	// ========================================
	// START WORKERS
	// ========================================

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

	// ========================================
	// START SCANNER
	// ========================================

	go func() {

		err := scanner.ScanDirectory(root, jobs)
		if err != nil {
			log.Println(err)
		}

	}()

	// ========================================
	// CLOSE RESULTS CHANNEL
	// ========================================

	go func() {
		wg.Wait()
		close(results)
	}()

	// ========================================
	// PROCESS RESULTS
	// ========================================

	for result := range results {

		fmt.Println("================================")
		fmt.Println("File:", result.Path)
		fmt.Println("Size:", result.Size, "bytes")
		fmt.Println("SHA256:", result.Checksum)

		// Protect shared state
		mu.Lock()

		// Store file
		allFiles = append(allFiles, result)

		// Duplicate grouping
		duplicates[result.Checksum] = append(
			duplicates[result.Checksum],
			result.Path,
		)

		mu.Unlock()
	}

	// ========================================
	// DUPLICATE REPORT
	// ========================================

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

	// ========================================
	// STORAGE ANALYTICS
	// ========================================

	fmt.Println("\n==============================")
	fmt.Println("📊 STORAGE ANALYTICS")
	fmt.Println("==============================")

	var totalSize int64

	for _, file := range allFiles {
		totalSize += file.Size
	}

	fmt.Println("Total Files:", len(allFiles))
	fmt.Println("Total Size:", totalSize, "bytes")

	// Largest files
	largest := analytics.GetLargestFiles(allFiles, 5)

	fmt.Println("\n🔥 Largest Files")

	for _, file := range largest {

		fmt.Println("--------------------------------")
		fmt.Println("File:", file.Path)
		fmt.Println("Size:", file.Size, "bytes")
	}
}