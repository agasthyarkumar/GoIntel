package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goi",
	Short: "Concurrent File Intelligence Tool",
	Long: `GoIntel is a concurrent file intelligence and security toolkit
built with Go using goroutines, channels, worker pools, and AES encryption.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}