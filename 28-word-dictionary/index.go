package main

import (
	"time"
	"os"
	"fmt"
	"strings"
	"bufio"
)

func createNewIndex() {
	start := time.Now()

	// create a new index file
	dat, err := os.Create(indexPath)
	logError("Error creating the new index file", err)
	defer dat.Close()

	// open the dictionary file
	file, err := os.Open(csvFilePath)
	logError("Error opening the dictionary file", err)
	defer file.Close()

	// initialize scanner
	scanner := bufio.NewScanner(file)
	var offset int64

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ",") // split the line
		_, err := dat.WriteString(fmt.Sprintf("%s %d\n", words[0], offset)) // store word space index
		logError("Error writing to the index file", err)

		offset += int64(len(line)) + int64(len("\r\n")) // include newline characters in offset
	}
	fmt.Printf("Created new index file successfully in %s\n\n", time.Since(start))
}
