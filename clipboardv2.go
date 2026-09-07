package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

var (
	savedClipboardv2 []byte
	clipboardSavedv2 bool
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
		savedClipboardv2 = clipboardData
		clipboardSavedv2 = true
	} else {
		fmt.Fprintln(os.Stderr, "Warning: No clipboard data to preserve.")
		savedClipboardv2 = nil
		clipboardSavedv2 = false
	}
}

func restoreClipboardV2() {
	if clipboardSavedv2 {
		writeClipboardV2(savedClipboardv2)
		clipboardSavedv2 = false // Prevent double restoration
		savedClipboardv2 = nil   // Clear saved clipboard data
	}
}
