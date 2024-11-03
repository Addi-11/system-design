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
	PacketName string
	AgentName string
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
	defer resp3.Body.Close()
	// Parse food name from the book response
	var packetResponse struct {
		PacketName string `json:"packet_name"`
	}
	if err := json.NewDecoder(resp3.Body).Decode(&packetResponse); err != nil {
		return nil, errors.New("failed to get packet name")
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
	defer resp4.Body.Close()
	// Parse agent name from the book response
	var agentResponse struct {
		AgentName string `json:"agent_name"`
	}
	if err := json.NewDecoder(resp4.Body).Decode(&agentResponse); err != nil {
		return nil, errors.New("failed to get delivery agent name")
	}

	return &Order{ID: orderID, PacketName: packetResponse.PacketName, AgentName: agentResponse.AgentName}, nil
}

func main(){
	var wg sync.WaitGroup

	for i:=0; i<10; i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			food_id := rand.Intn(2)+1 // randomly choose burger or pizza to order
			order, err := PlaceOrder(food_id)
			if err != nil{
				fmt.Println("order not placed:", err.Error())
			}else {
				fmt.Printf("order_id = %s\norder placed for %s, delivery agent %s\n\n", order.ID, order.PacketName, order.AgentName)
			}
		}()
	}
	wg.Wait()
	fmt.Println("all orders completed")
}