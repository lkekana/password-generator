package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

var (
	savedClipboard []byte
	clipboardSaved bool
)

func readClipboardV2() []byte {
	clipboardData, err := clipboard.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Warning: Failed to read clipboard:", err)
		return nil
	}
	return []byte(clipboardData)
}

func writeClipboardV2(data []byte) {
	err := clipboard.WriteAll(string(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Warning: Failed to write to clipboard:", err)
	}
}

func preserveClipboardV2() {
	clipboardData := readClipboardV2()
	if clipboardData != nil && len(clipboardData) > 0 {
		savedClipboard = clipboardData
		clipboardSaved = true
	} else {
		fmt.Fprintln(os.Stderr, "Warning: No clipboard data to preserve.")
		savedClipboard = nil
		clipboardSaved = false
	}
}

func restoreClipboardV2() {
	if clipboardSaved {
		writeClipboardV2(savedClipboard)
		clipboardSaved = false // Prevent double restoration
		savedClipboard = nil   // Clear saved clipboard data
	}
}
