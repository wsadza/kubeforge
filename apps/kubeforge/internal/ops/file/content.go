// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
//
// ############################################################

package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Define the mode type
type ReadMode int

const (
	// Define custom modes for reading the file
	ReadModeByte   ReadMode = iota // Mode to read file as []byte
	ReadModeString                 // Mode to read file as string
	ReadModeSlice                  // Mode to read file as []string (lines)
)

// Read function reads the file based on the mode and returns the appropriate type
func Read(filePath string, readMode ...ReadMode) (interface{}, error) {
	// Set default mode to ReadModeByte if no mode is passed
	var mode ReadMode

	if len(readMode) == 0 {
		mode = ReadModeByte
	} else {
		mode = readMode[0] // Use the first provided mode
	}

	var result interface{}

	// Open the file in read-only mode
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	// Switch based on the read mode
	switch mode {

	case ReadModeString:
		// Read the entire file as a string
		content, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("Error reading file as string: %v", err)
		}
		result = string(content) // Convert []byte to string

	case ReadModeSlice:
		// Read the file as []string (lines)
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("Error reading file as lines: %v", err)
		}
		result = lines

	default:
		// If an unsupported mode is provided, fall back to reading as []byte (ReadModeByte)
		result, err = io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("Error reading file as byte array: %v", err)
		}
	}

	return result, nil
}

// ------------------------------------------------------------

// Define the mode type
type WriteMode int

const (
	// Define custom modes for reading the file
  WriteModeOverride ReadMode = iota // Mode to override content of file 
	WriteModeAppend                   // Mode to append content into given line 
	WriteModeChange                   // Mode to change content from given line 
)

func Write(filePath string, data interface{}, writeMode ...WriteMode) error { 

  return nil
}

// ------------------------------------------------------------

// Find searches for a string in a file and returns the line number where the
// string is located. 
// If the string is not found, it returns -1 and no error.
func Find(filePath, searchString string) (int, error) {
	// Open the file in read-only mode
	file, err := os.Open(filePath)
	if err != nil {
		return -1, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// Check if the line contains the search string
		if strings.Contains(line, searchString) {
			return lineNumber, nil
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return -1, fmt.Errorf("error reading file: %w", err)
	}

	// String not found
	return -1, nil
}

// ------------------------------------------------------------

func Expand(fileBody string) string {
  return os.ExpandEnv(fileBody)
}
