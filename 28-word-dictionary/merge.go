package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

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
			writeDictionaryEntry(dict, newDict, newIndex, dictWord, dictOffset, &offset)


			// move dictoffset pointer
			if idxScanner.Scan() {
				dictWord, dictOffset = parseIndexEntry(idxScanner.Text())
			} else {
				dictWord = ""
			}
		} else if dictWord == "" || strings.Compare(chglogWord, dictWord) <= 0 {
			if strings.Compare(chglogWord, dictWord) == 0{
				// move dictoffset pointer
				if idxScanner.Scan() {
					dictWord, dictOffset = parseIndexEntry(idxScanner.Text())
				} else {
					dictWord = ""
				}
			}
			
			writeChangelogEntry(newDict, newIndex, chglogWord, chglogMeaning, &offset)
			// move changelog pointer
			if chglogScanner.Scan() {
				chglogWord, chglogMeaning = parseChangelogEntry(chglogScanner.Text())
			} else {
				chglogWord = ""
			}	
		}
	}

	fmt.Println("Merge completed: New dictionary and index created.")
}

