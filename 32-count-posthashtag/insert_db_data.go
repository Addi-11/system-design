package main

import (
	"fmt"
	"os"
	"math/rand"
)

func insertDBData() {
	file, err := os.Create("setup_db.sql")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()


	file.WriteString("CREATE DATABASE IF NOT EXISTS test_db;\n")
	file.WriteString("USE test_db;\n")
	file.WriteString("CREATE TABLE hashtags (\n")
	file.WriteString("    hashtag_id VARCHAR(50) PRIMARY KEY,\n")
	file.WriteString("    count INT\n")
	file.WriteString(");\n\n")

	// Write insert statements
	// num of uniques hastags
	file.WriteString("INSERT INTO hashtags (hashtag_id, count) VALUES\n")
	for i := 2001; i <= 5000; i++ {
		cnt := rand.Intn(10000)
		if i != numHashtag {
			file.WriteString(fmt.Sprintf("('#tag%d', %d),\n", i, cnt))
		} else {
			file.WriteString(fmt.Sprintf("('#tag%d', %d);\n", i, cnt))
		}
	}

	fmt.Println("SQL file 'setup_db.sql' generated successfully!")
}
