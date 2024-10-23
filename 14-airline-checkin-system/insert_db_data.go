package main

import (
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v6"
)

func insert_seats(){
	file, err := os.OpenFile("set_db.sql", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\nINSERT INTO seats (seat_no) VALUES\n")

	seatLetters := []string{"A", "B", "C", "D", "E", "F"}
	seatRows := 20

	for row := 1; row <= seatRows; row++{
		for _, letter := range seatLetters{
			seat := fmt.Sprintf("('%d%s')", row, letter)
			if row == seatRows && letter == "F" {
				_, err = file.WriteString(seat + ";\n")
			} else {
				_, err = file.WriteString(seat + ", ")
			}
			if err != nil {
				fmt.Println("Error writing to file", err)
				return
			}
		}
	}
	fmt.Println("Appended seats to set_db.sql")
}

func insert_users(){
	file, err := os.OpenFile("set_db.sql", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\nINSERT INTO users (name) VALUES\n")

	for i:=0; i<120; i++{
		name := gofakeit.Name()
		entry := fmt.Sprintf("('%s')", name)

		if i == 119{
			_, err = file.WriteString(entry + ";\n")
		} else {
			_, err = file.WriteString(entry + ", ")
		}
		if err != nil{
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	fmt.Println("Appended users to set_db.sql")
}