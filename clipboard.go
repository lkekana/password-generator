package main

import (
	"fmt"
	"os"

	"golang.design/x/clipboard"
)

var initRun bool = false

var (
	savedClipboard []byte
	savedFormat    clipboard.Format
	clipboardSaved bool
)

func initClipboard() {
	if initRun {
		return
	}

	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Warning: Failed to initialize clipboard:", err)
		panic(err)
	}
	initRun = true
}

func readClipboard() ([]byte, clipboard.Format) {
	initClipboard()

	// clipboard.Read returns nil if the format is wrong, so check if Image first
	clipboardData := clipboard.Read(clipboard.FmtImage)
	if clipboardData != nil {
		return clipboardData, clipboard.FmtImage
	}

	return clipboard.Read(clipboard.FmtText), clipboard.FmtText
}

func writeClipboard(data []byte, format clipboard.Format) {
	initClipboard()
	clipboard.Write(format, data)
}

func preserveClipboard() {
	initClipboard()
	clipboardData, format := readClipboard()
	if clipboardData != nil && len(clipboardData) > 0 {
		savedClipboard = clipboardData
		savedFormat = format
		clipboardSaved = true
	} else {
		fmt.Fprintln(os.Stderr, "Warning: No clipboard data to preserve.")
		savedClipboard = nil
		clipboardSaved = false
	}
}

func restoreClipboard() {
	if clipboardSaved {
		writeClipboard(savedClipboard, savedFormat)
		clipboardSaved = false // Prevent double restoration
		savedClipboard = nil   // Clear saved clipboard data
	}
}
