package main

import (
	"fmt"
	"os"
	"path/filepath"
	"unsafe"

	"golang.design/x/clipboard"
)

// it's 8 bytes (size of int) on my computer but just in case
func getSizeOfClipboardFormat() int {
	sizeTxt := unsafe.Sizeof(clipboard.FmtText)
	sizeImg := unsafe.Sizeof(clipboard.FmtImage)
	return max(int(sizeTxt), int(sizeImg))
}

func getAppDir() string {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting config directory:", err)
		return ""
	}

	appDir := filepath.Join(baseDir, appName)

	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		fmt.Println("Error creating app directory:", err)
		return ""
	}

	return appDir
}

func readClipboard() []byte {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	return clipboard.Read(clipboard.FmtText)
}

func writeClipboard(data []byte, format clipboard.Format) {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(format, data)
}

func preserveClipboard(data []byte, format clipboard.Format) {
	current := readClipboard()

	appDir := getAppDir()
	if appDir == "" {
		fmt.Println("Failed to get or create app directory.")
		return
	}

	file, err := os.Create(filepath.Join(appDir, ".clipboard"))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	// write the format with as many bytes as the size of the clipboard format (8 bytes on my computer)
	// then write the current clipboard contents
	formatBytes := make([]byte, getSizeOfClipboardFormat())
	formatBytes = append(formatBytes, byte(format))
	fileContents := append(formatBytes, current...)

	_, err = file.Write(fileContents)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	err = file.Sync()
	if err != nil {
		fmt.Println("Error syncing file:", err)
		return
	}
}

func restoreClipboard() {
	appDir := getAppDir()
	if appDir == "" {
		fmt.Println("Failed to get or create app directory.")
		return
	}

	clipboardFormatSize := getSizeOfClipboardFormat()

	file, err := os.Open(filepath.Join(appDir, ".clipboard"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fileSize := fileInfo.Size()
	if fileSize < int64(clipboardFormatSize) {
		fmt.Println("File is too small to contain clipboard data.")
		return
	}

	formatBytes := make([]byte, clipboardFormatSize)
	_, err = file.Read(formatBytes)
	if err != nil {
		fmt.Println("Error reading format from file:", err)
		return
	}

	format := clipboard.Format(formatBytes[len(formatBytes)-1])

	data := make([]byte, fileSize-int64(clipboardFormatSize))
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading data from file:", err)
		return
	}

	writeClipboard(data, format)
}
