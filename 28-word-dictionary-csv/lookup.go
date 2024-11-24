package main

import (
	"time"
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

func writeDictionaryEntry(dict *os.File, newDict, newIndex *os.File, word string, dictOffset int64, offset *int64) {
	dict.Seek(dictOffset, 0)
	line, _ := bufio.NewReader(dict).ReadString('\n')
	newDict.WriteString(line)
	newIndex.WriteString(fmt.Sprintf("%s %d\n", word, *offset))
	*offset += int64(len(line))
}

func wordLookup(word string) string {
	start := time.Now()
	// open and read the dat file - to search for word index
	dat, err := os.Open(indexPath)
	logError("Error opening index file", err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	var index int64 = -1

	for scanner.Scan() {
		line := scanner.Text()
		scannedWord := strings.Split(line, " ")

		if word == scannedWord[0] {
			index, _ = strconv.ParseInt(scannedWord[1], 10, 64)
			fmt.Printf("Found word %s at byte offset %d\n", scannedWord[0], index)
			break
		}
	}

	// if word not found
	if index == -1 {
		fmt.Printf("Word %s not found in dictionary\n", word)
		return ""
	}

	// open the csv file and directly read from the byte offset
	file, err := os.Open(csvFilePath)
	logError("Error opening the dictionary file", err)
	defer file.Close()

	// read line from the offset
	file.Seek(index, 0) // Seek to the offset
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n') // read until newline
	content := strings.Split(line, ",")

	// word, meaning
	if len(content) < 2 {
		fmt.Printf("Error: Malformed line at offset %d\n", index)
		return ""
	}
	fmt.Printf("%s : %s\nTime taken: %s\n", content[0], content[1], time.Since(start))

	return content[1]
}
