package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"

	"golang.design/x/clipboard"
)

func getAppDir() string {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting config directory:", err)
		return ""
	}

	appDir := filepath.Join(baseDir, appName)

	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating app directory:", err)
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
		fmt.Fprintln(os.Stderr, "Failed to get or create app directory.")
		return
	}

	file, err := os.Create(filepath.Join(appDir, ".clipboard"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating file:", err)
		return
	}

	defer file.Close()

	// let Go handle the size of the clipboard format by writing it as a binary value
	err = binary.Write(file, binary.LittleEndian, format)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing format to file:", err)
		return
	}

	_, err = file.Write(current)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing to file:", err)
		return
	}

	err = file.Sync()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error syncing file:", err)
		return
	}
}

func restoreClipboard() {
	appDir := getAppDir()
	if appDir == "" {
		fmt.Fprintln(os.Stderr, "Failed to get or create app directory.")
		return
	}

	file, err := os.Open(filepath.Join(appDir, ".clipboard"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "No clipboard backup found to restore.")
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting file info:", err)
		return
	}
	fileSize := fileInfo.Size()

	var format clipboard.Format
	err = binary.Read(file, binary.LittleEndian, &format)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading clipboard format from backup file:", err)
		return
	}

	data := make([]byte, fileSize-int64(binary.Size(format)))
	_, err = file.Read(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading data from backup file:", err)
		return
	}

	writeClipboard(data, format)

	// delete the file
	err = os.Remove(filepath.Join(appDir, ".clipboard"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error deleting backup file:", err)
		return
	}
}
