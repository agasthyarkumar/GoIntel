package cmd

import (
	"fmt"
	"log"

	"gointel/internal/compression"

	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:     "compress [path]",
	Aliases: []string{"c"},
	Short:   "Compress file or folder into ZIP",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		path := args[0]

		fmt.Println("📦 Compressing:", path)

		err := compression.Compress(path)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("✅ Compression complete")
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
}