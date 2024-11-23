package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	csvFilePath = "dictionary.csv"
	indexPath = "index.dat"
)

func createNewIndex(){
	start := time.Now()

	// create a new index file
	dat, err := os.Create(indexPath)
	if err != nil{
		log.Fatal("Error creating the new index file", err)
	}
	defer dat.Close()

	// open the dictionary file
	file, err := os.Open(csvFilePath)
	if err != nil{
		log.Fatal("Error opening the dictionary file", err)
	}
	defer file.Close()

	// we loop thorugh the dictionary file, to create the byte offset index
	// initialize scanner
	scanner := bufio.NewScanner(file)
	var offset int64

	for scanner.Scan(){
		line := scanner.Text()
		words := strings.Split(line, ",") // split the line 
		_, err := dat.WriteString(fmt.Sprintf("%s %d\n", words[0], offset)) // store word space index
		if err != nil{
			log.Fatal("Error writing to the index file", err)
		}
		offset += int64(len(line)) + int64(len("\r\n")) // include newline chrct in offset in WINDOWS
	}
	fmt.Printf("Created new index file successfully in %s\n\n", time.Since(start))
}

func createNewCSV(){
	file, err := os.Create(csvFilePath)
	if err != nil{
		log.Fatal("Error creating new dictionary file", err)
	}

	defer file.Close()

	// new csv writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

}

func wordLookup(word string) string{
	start := time.Now()
	// open and read the dat file - to search for word index
	dat, err := os.Open(indexPath)
	if err != nil{
		log.Fatal("Error opening index file", err) 
	}
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	var index int64 = -1

	for scanner.Scan(){
		line := scanner.Text()
		scannedWord := strings.Split(line, " ")

		if len(scannedWord) != 2 {
			log.Fatal("Malformed index file entry:", line)
		}
		if word == scannedWord[0]{
			index, _ = strconv.ParseInt(scannedWord[1], 10, 64)
			fmt.Printf("Found word %s at byte offset %d\n", scannedWord[0], index)
			break
		}
	}

	// if word not found
	if index == -1 {
		fmt.Printf("Word %s not fond in dictionary", word)
		return ""
	}

	// open the csv file and directly read from the byte offset
	file, err := os.Open(csvFilePath)
	if err != nil{
		log.Fatal("Error opening the index file", err)
	}
	defer file.Close()
	
	// read line from the offset
	file.Seek(index, 0) // Seek to the offset
	// currentOffset, _ := file.Seek(0, os.SEEK_CUR)
	// fmt.Printf("Debug: Current file pointer is at offset %d\n", currentOffset)


	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\r') // WINDOWS \r
	content := strings.Split(line, ",")

	// word, meaning
	fmt.Printf("%s : %s\nTime taken: %s\n\n", content[0], content[1], time.Since(start))

	return content[1]
}

func addWord(newWord string){

}

// to update word - first scan the index to see if word is present
// if not present we append to the dictionary and index in sorted order - merge sort
// if present we create a change log and merge old index and old dictionary - APP 2, 
// if present we create a new index and new dictionary and update meta.json which has these details - APP3

func main(){
	createNewIndex()
	wordLookup("Apple")
	wordLookup("Dog")


}