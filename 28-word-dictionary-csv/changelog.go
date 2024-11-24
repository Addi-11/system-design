package main

import (
	"os"
	"fmt"
)

func writeChangelogEntry(newDict, newIndex *os.File, word, meaning string, offset *int64) {
	line := fmt.Sprintf("%s,%s\n", word, meaning)
	newDict.WriteString(line)
	newIndex.WriteString(fmt.Sprintf("%s %d\n", word, *offset))
	*offset += int64(len(line))
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