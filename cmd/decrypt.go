package cmd

import (
	"fmt"
	"log"

	"gointel/internal/crypto"

	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:     "decrypt [file]",
	Aliases: []string{"d"},
	Short:   "Decrypt AES encrypted file",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		file := args[0]

		fmt.Println("🔓 Decrypting:", file)

		err := crypto.DecryptFile(file)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("✅ Decryption complete")
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}