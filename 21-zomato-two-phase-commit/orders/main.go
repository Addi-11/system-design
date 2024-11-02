// main order service which will place the user's order
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type Order struct{
	ID string
}

func PlaceOrder(foodID int) (*Order, error){

	// reserve food
	body, _ := json.Marshal(map[string]interface{}{
		"food_id": foodID,
	})
	reqBody := bytes.NewBuffer(body)
	resp1, err := http.Post("http://localhost:8081/food/reserve", "application/json", reqBody)
	if err != nil || resp1.StatusCode != 200 {
		return nil, errors.New("food not available")
	}

	//reserve agent
	resp2, err := http.Post("http://localhost:8082/delivery/agent/reserve", "applciation/json", nil)
	if err != nil || resp2.StatusCode != 200 {
		return nil, errors.New("delivery agent not available")
	}

	// create a new order id
	orderID := uuid.New().String()

	// book food
	body, _ = json.Marshal(map[string]interface{}{
		"food_id": foodID,
		"order_id": orderID,
	})
	reqBody = bytes.NewBuffer(body)
	resp3, err := http.Post("http://localhost:8081/food/book", "application/json", reqBody)
	if err != nil || resp3.StatusCode != 200 {
		return nil, errors.New("could not assign food to the order")
	}

	// book agent
	body, _ = json.Marshal(map[string]interface{}{
		"order_id": orderID,
	})
	reqBody = bytes.NewBuffer(body)
	resp4, err := http.Post("http://localhost:8082/delivery/agent/book", "application/json", reqBody)
	if err != nil || resp4.StatusCode != 200 {
		return nil, errors.New("could not assign delivery agent to the order")
	}

	return &Order{ID: orderID}, nil
}

func main(){
	var wg sync.WaitGroup

	for i:=0; i<10; i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			food_id := rand.Intn(2)+1 // randomly choose burger or pizza to order
			order_id, err := PlaceOrder(food_id)
			if err != nil{
				fmt.Println("order not placed:", err.Error())
			}else {
				fmt.Println("order placed: ", order_id)
			}
		}()
	}
	wg.Wait()
	fmt.Println("all orders completed")
}