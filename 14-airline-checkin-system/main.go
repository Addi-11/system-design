package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type User struct{
	ID int
	Name string
	Seat string
}

func init(){
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"
	dbName := "airline_checkin"

	// connection string to the DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	// open connection to the DB
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil{
		panic(err)
	}

	// test DB connection
	err = DB.Ping()
	if err != nil{
		log.Fatal("Error pinging DB :", err)
	}
}

func assign_seat(user *User) error{
	// begin transaction
	txn, _ := DB.Begin()

	// get empty seat
	row := txn.QueryRow("SELECT seat_no from seats WHERE user_id IS null ORDER BY seat_no LIMIT 1 FOR UPDATE SKIP LOCKED")
	row.Scan(&user.Seat)

	// update the empty seat
	_, err := txn.Exec("UPDATE seats SET user_id=? WHERE seat_no=?", user.ID, user.Seat) 
	if err != nil{
		return fmt.Errorf("error updating seat: %v", err)
	}
	err = txn.Commit()
	if err != nil{
		return fmt.Errorf("error committing transaction: %v", err)
	}
	return nil
}

func get_users() []User{
	users := make([]User, 0)
	rows, err := DB.Query("SELECT id, name FROM users ORDER BY id;")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil{
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func print_airline(){
	fmt.Print("\n\nFlight Seating Chart:\n")
	seatLetters := []string{"A", "B", "C", "D", "E", "F"}
	seatRows := 20

	for _, letter := range(seatLetters){
		for row := 1; row <= seatRows; row++{

			seat_no := fmt.Sprintf("%d%s", row, letter)
			var user_id *int
			err := DB.QueryRow("SELECT user_id from seats where seat_no = ?", seat_no).Scan(&user_id)
			
			if err != nil{
				log.Println("Error querying the seat: ", seat_no," ", err)
				continue
			}

			if user_id != nil{
				fmt.Print("x") // seat occupied
			}else{
				fmt.Print(".") // seat unoccupied
			}
		}
		fmt.Print("\n")
		if letter == "C"{
			fmt.Print("\n")
		}
	}
}

func main(){
	users := get_users()
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(len(users))

	for ix := range users {
		go func(user *User){
			err := assign_seat(user)
			if err != nil{
				fmt.Printf("Could not assign seat to %s\n", user.Name)
			}else{
				fmt.Printf("Assigned %s, seat %s\n", user.Name, user.Seat)
			}
			wg.Done()
		}(&users[ix])
	}

	wg.Wait()
	fmt.Println("\nTime taken to assign seats:", time.Since(start))
	print_airline()
}