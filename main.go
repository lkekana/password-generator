package main

import (
	"fmt"
	"os"
	"time"
	"unsafe"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var appName = "pwg"
var length int
var includeUppercase bool
var includeLowercase bool
var includeNumbers bool
var includeSpecialChars bool
var count int
var debug bool
var printWithNewline bool

func main() {
	rootCmd := &cobra.Command{
		Use:   "pwg",
		Short: "A simple password generator CLI",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if !debug {
				if count == 1 && !printWithNewline {
					// avoid printing a newline when only one password is generated
					password, err := generatePassword(length, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars)
					if err != nil {
						fmt.Fprintln(os.Stderr, "Error generating password:", err)
						return
					}
					os.Stdout.Write(password)
				} else {
					for i := 0; i < count; i++ {
						password, err := generatePassword(length, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars)
						if err != nil {
							fmt.Fprintf(os.Stderr, "Error generating password %d: %v\n", i+1, err)
							continue
						}
						os.Stdout.Write(password)
						os.Stdout.Write([]byte("\n"))
					}
				}
			} else {
				err := clipboard.Init()
				if err != nil {
					panic(err)
				}

				sizeTxt := unsafe.Sizeof(clipboard.FmtText)
				sizeImg := unsafe.Sizeof(clipboard.FmtImage)
				color.Yellow("Debug mode enabled.")

				fmt.Println("\n=== GENERATOR CONFIGURATION ===")
				fmt.Println("Number of passwords to generate:", color.MagentaString("%d", count))
				fmt.Println("Length of each password:", color.MagentaString("%d", length))
				fmt.Println("Include uppercase:", color.YellowString("%t", includeUppercase))
				fmt.Println("Include lowercase:", color.YellowString("%t", includeLowercase))
				fmt.Println("Include numbers:", color.YellowString("%t", includeNumbers))
				fmt.Println("Include special characters:", color.YellowString("%t", includeSpecialChars))
				fmt.Println("===============================")

				fmt.Println("\n=== CLIPBOARD INFORMATION ===")
				fmt.Println("Size of clipboard.FmtText:", color.MagentaString("%d bytes", sizeTxt))
				fmt.Println("Size of clipboard.FmtImage:", color.MagentaString("%d bytes", sizeImg))
				clip := clipboard.Read(clipboard.FmtImage)
				if clip == nil {
					fmt.Println("Clipboard does not contain image data.")
					clip = clipboard.Read(clipboard.FmtText)
					if clip == nil {
						fmt.Println("Clipboard does not contain text data either.")
					} else {
						fmt.Println("Clipboard contains text data of size:", color.MagentaString("%d bytes", len(clip)))
						fmt.Println("Clipboard content:", string(clip))
					}
				} else {
					fmt.Println("Clipboard contains image data of size:", color.MagentaString("%d bytes", len(clip)))
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
				}
				initElapsed := time.Since(initStart)
				fmt.Println("Total execution time for", count, "passwords:", color.GreenString("%s", initElapsed))
				fmt.Println()
			}
		},
	}

	rootCmd.Flags().IntVarP(&length, "length", "l", 16, "Length of the password")
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
