package main

import (
	"strings"
	"strconv"
	"log"
)

func logError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func parseIndexEntry(entry string) (string, int64) {
	parts := strings.Split(entry, " ")
	offset, _ := strconv.ParseInt(parts[1], 10, 64)
	return parts[0], offset
}

func parseChangelogEntry(entry string) (string, string) {
	parts := strings.Split(entry, ",")
	return parts[0], parts[1]
}