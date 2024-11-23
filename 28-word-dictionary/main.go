
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	csvFilePath    = "dictionary.csv"
	indexPath      = "index.dat"
	changeLogPath  = "changelog.csv"
	newCsvFilePath = "new_dictionary.csv"
	newIndexPath   = "new_index.dat"
)

func logError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

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
	fmt.Printf("%s : %s\nTime taken: %s\n\n", content[0], content[1], time.Since(start))

	return content[1]
}

func merge() {
	// Open dictionary
	dict, err := os.Open(csvFilePath)
	logError("Error opening dictionary", err)
	defer dict.Close()

	// Open dictionary index
	idx, err := os.Open(indexPath)
	logError("Error opening index file", err)
	defer idx.Close()

	// Open changelog
	chglog, err := os.Open(changeLogPath)
	logError("Error opening changelog file", err)
	defer chglog.Close()

	// Create files for new dictionary
	newDict, err := os.Create(newCsvFilePath)
	logError("Error creating new dictionary", err)
	defer newDict.Close()

	// Create new index file for new dictionary
	newIndex, err := os.Create(newIndexPath)
	logError("Error creating new index", err)
	defer newIndex.Close()

	// read files
	var dictWord, chglogWord, chglogMeaning string
	var dictOffset, offset int64

	idxScanner := bufio.NewScanner(idx)
	chglogScanner := bufio.NewScanner(chglog)

	idxScanner.Scan()
	chglogScanner.Scan()

	chglogParts := strings.Split(chglogScanner.Text(), ",")
	chglogWord = chglogParts[0]
	chglogMeaning = chglogParts[1]
	dictParts := strings.Split(idxScanner.Text(), " ")
	dictWord = dictParts[0]
	dictOffset, _ = strconv.ParseInt(dictParts[1], 10, 64)

	// use pointers to merge
	for dictWord != "" || chglogWord != "" {
		if chglogWord == "" || (dictWord != "" && strings.Compare(dictWord, chglogWord) < 0) {
			dict.Seek(dictOffset, 0)
			dictLine, _ := bufio.NewReader(dict).ReadString('\n')

			newDict.WriteString(dictLine)
			newIndex.WriteString(fmt.Sprintf("%s %d\n", dictWord, offset))
			offset += int64(len(dictLine))

			// move dictoffset pointer
			if idxScanner.Scan() {
				dictParts := strings.Split(idxScanner.Text(), " ")
				dictWord = dictParts[0]
				dictOffset, _ = strconv.ParseInt(dictParts[1], 10, 64)
			} else {
				dictWord = ""
			}
		} else if dictWord == "" || strings.Compare(chglogWord, dictWord) <= 0 { // = in case the word exists in both
			newDict.WriteString(fmt.Sprintf("%s,%s\n", chglogWord, chglogMeaning))
			newIndex.WriteString(fmt.Sprintf("%s %d\n", chglogWord, offset))
			offset += int64(len(chglogWord) + len(chglogMeaning) + 1) + int64(len("\n"))

			// move changelog pointer
			if chglogScanner.Scan() {
				chglogParts := strings.Split(chglogScanner.Text(), ",")
				chglogWord = chglogParts[0]
				chglogMeaning = chglogParts[1]
			} else {
				chglogWord = ""
			}
		}
	}

	fmt.Println("Merge completed: New dictionary and index created.")
}

func addWordToChangelog(word, meaning string) {
	file, err := os.OpenFile(changeLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // creates when file not present
	logError("Error creating the changelog file", err)
	defer file.Close()

	// append word to the changelog
	_, err = file.WriteString(fmt.Sprintf("%s,%s\n", word, meaning))
	logError("Error appending to the changelog", err)

	fmt.Printf("Word added to changelog: %s\n", word)
}

func addWord(word, meaning string) {
	addWordToChangelog(word, meaning)
	merge() // merge dictionary and changelog

	os.Rename(newCsvFilePath, csvFilePath)
	os.Rename(newIndexPath, indexPath)
	os.Remove(changeLogPath)
}

func main() {
	createNewIndex()
	wordLookup("Apple")
	wordLookup("Dog")

	addWord("Banana", "Yellow fruit often eaten by monkeys and humans alike")
	addWord("Elephant", "Large mammal with a trunk and grey colored")
	addWord("Funny", "A trait many people lack.")
	addWord("Dog", "Bow bow bow wow wow bark")

	wordLookup("Elephant")
	wordLookup("Orange")
	wordLookup("Dog")
}