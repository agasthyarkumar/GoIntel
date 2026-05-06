package cmd

import (
	"fmt"
	"log"

	"gointel/internal/crypto"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:     "encrypt [file]",
	Aliases: []string{"e"},
	Short:   "Encrypt a file using AES-256",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		file := args[0]

		fmt.Println("🔐 Encrypting:", file)

		err := crypto.EncryptFile(file)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("✅ Encryption complete")
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}