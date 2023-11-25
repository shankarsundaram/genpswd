/*
Copyright © 2023 Shankar Sundaram shankar.sundaram@outlook.com
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Random Passwords",
	Long: `Generate Random Passwords with customizable options.
	Example:
	genpswd generate -n 32 -sdul`,

	Run: generatePassword,
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	symbols, _ := cmd.Flags().GetBool("symbols")
	digits, _ := cmd.Flags().GetBool("digits")
	uppercase, _ := cmd.Flags().GetBool("uppercase")
	lowercase, _ := cmd.Flags().GetBool("lowercase")
	charset := ""
	if symbols {
		charset += "!@#$%^&*()"
	}
	if digits {
		charset += "0123456789"
	}
	if uppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if lowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if len(charset) == 0 {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()"
	}
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println(string(password))
}
func init() {
	rootCmd.AddCommand(passwordCmd)
	passwordCmd.Flags().IntP("length", "n", 12, "Length of the password")
	passwordCmd.Flags().BoolP("symbols", "s", false, "Include symbols in the password")
	passwordCmd.Flags().BoolP("digits", "d", false, "Include digits in the password")
	passwordCmd.Flags().BoolP("uppercase", "u", false, "Include uppercase letters in the password")
	passwordCmd.Flags().BoolP("lowercase", "l", false, "Include lowercase letters in the password")
}
