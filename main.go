package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var appName = "password-generator"
var length int
var includeUppercase bool
var includeLowercase bool
var includeNumbers bool
var includeSpecialChars bool
var count int
var debug bool
var printWithNewline bool
var copyToClipboard bool

func main() {
	rootCmd := &cobra.Command{
		Use:   "password-generator",
		Short: "A simple password generator CLI",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if length <= 0 {
				fmt.Fprintln(os.Stderr, color.RedString("Error: Password length must be greater than 0"))
				return
			}

			if !includeUppercase && !includeLowercase && !includeNumbers && !includeSpecialChars {
				fmt.Fprintln(os.Stderr, color.RedString("Error: At least one character set must be selected"))
				return
			}

			if count <= 0 {
				fmt.Fprintln(os.Stderr, color.RedString("Error: Count must be greater than 0"))
				return
			}

			if copyToClipboard && count > 1 {
				color.Yellow("Warning: Copying multiple passwords to clipboard is not supported. Only 1 password will be generated and copied to clipboard.")
				count = 1
				printWithNewline = false
			}

			if !debug {
				if count == 1 && !printWithNewline {
					// avoid printing a newline when only one password is generated
					password, err := generatePassword(length, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars)
					if err != nil {
						fmt.Fprintln(os.Stderr, "Error generating password:", err)
						return
					}
					if copyToClipboard {
						// preserveClipboard()
						writeClipboardV2(password)
						zeroOutPassword(password)

						fmt.Println(color.GreenString("Password copied to clipboard."))
						// fmt.Println(color.YellowString("Note: Clipboard will be cleared in 10 seconds."))

						// time.Sleep(10 * time.Second)

						// restoreClipboard()
						return
					} else {
						os.Stdout.Write(password)
						zeroOutPassword(password)
					}
				} else {
					for i := 0; i < count; i++ {
						password, err := generatePassword(length, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars)
						if err != nil {
							fmt.Fprintf(os.Stderr, "Error generating password %d: %v\n", i+1, err)
							continue
						}
						os.Stdout.Write(password)
						zeroOutPassword(password)
						os.Stdout.Write([]byte("\n"))
					}
				}
			} else {
				color.Yellow("Debug mode enabled.")
				if copyToClipboard {
					color.Yellow("Warning: Copying to clipboard is enabled, but debug mode is also enabled. Clipboard operations will be logged but not executed.")
				}

				fmt.Println("\n=== GENERATOR CONFIGURATION ===")
				fmt.Println("Number of passwords to generate:", color.MagentaString("%d", count))
				fmt.Println("Length of each password:", color.MagentaString("%d", length))
				fmt.Println("Include uppercase:", color.YellowString("%t", includeUppercase))
				fmt.Println("Include lowercase:", color.YellowString("%t", includeLowercase))
				fmt.Println("Include numbers:", color.YellowString("%t", includeNumbers))
				fmt.Println("Include special characters:", color.YellowString("%t", includeSpecialChars))
				fmt.Println("===============================")

				fmt.Println("\n=== CLIPBOARD INFORMATION ===")
				clip := readClipboardV2()
				if clip == nil || len(clip) == 0 {
					fmt.Println("Clipboard does not contain any data.")
				} else {
					fmt.Println("Clipboard contains data of size:", color.MagentaString("%d bytes", len(clip)))
					// fmt.Println("Clipboard content:", string(clip))
				}
				fmt.Println("=============================")
				fmt.Println()

				initStart := time.Now()
				for i := 0; i < count; i++ {
					start := time.Now()
					password, err := generatePassword(length, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error generating password %d: %v\n", i+1, err)
						continue
					}
					elapsed := time.Since(start)
					fmt.Printf("Password %d: ", i+1)
					os.Stdout.Write(password)
					fmt.Printf(" (Execution took %s)\n", elapsed.String())
					zeroOutPassword(password)
				}
				initElapsed := time.Since(initStart)
				fmt.Println("Total execution time for", count, "passwords:", color.GreenString("%s", initElapsed))
				fmt.Println()
			}
		},
	}

	rootCmd.Flags().IntVarP(&length, "length", "l", 16, "Length of the password")
	// rootCmd.Flags().BoolVarP(&copyToClipboard, "copy", "y", false, "Copy to clipboard and clear after 10 seconds")
	rootCmd.Flags().BoolVarP(&copyToClipboard, "copy", "y", false, "Copy to clipboard (only works when generating a single password)")
	rootCmd.Flags().BoolVar(&includeUppercase, "upper", true, "Include uppercase letters")
	rootCmd.Flags().BoolVar(&includeLowercase, "lower", true, "Include lowercase letters")
	rootCmd.Flags().BoolVar(&includeNumbers, "num", true, "Include numbers")
	rootCmd.Flags().BoolVar(&includeSpecialChars, "special", false, "Include special characters")
	rootCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of passwords to generate")
	rootCmd.Flags().BoolVar(&printWithNewline, "newline", false, "Print a newline after generating a single password (useful for piping output & does not apply when generating multiple passwords)")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error executing command:", err)
		os.Exit(1)
	}
}
