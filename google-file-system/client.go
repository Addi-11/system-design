package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	MasterAddress string
}

func (c *Client) WriteFile(filePath, data string) {
	resp, err := http.Get(fmt.Sprintf("http://%s/write_file?filePath=%s&data=%s", c.MasterAddress, filePath, data))
	if err != nil {
		log.Fatalf("Error communicating with master: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	chunkServer := result["chunkServer"]
	chunkID := result["chunkID"]

	_, err = http.Get(fmt.Sprintf("http://%s/write_chunk?chunkID=%s&data=%s", chunkServer, chunkID, data))
	if err != nil {
		log.Fatalf("Error communicating with chunk server: %v", err)
	}
	fmt.Println("File written successfully.")
}

func (c *Client) ReadFile(filePath string) {
	resp, err := http.Get(fmt.Sprintf("http://%s/read_file?filePath=%s", c.MasterAddress, filePath))
	if err != nil {
		log.Fatalf("Error communicating with master: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	chunkServer := result["chunkServer"]
	chunkID := result["chunkID"]

	chunkResp, err := http.Get(fmt.Sprintf("http://%s/read_chunk?chunkID=%s", chunkServer, chunkID))
	if err != nil {
		log.Fatalf("Error communicating with chunk server: %v", err)
	}
	defer chunkResp.Body.Close()

	data, _ := ioutil.ReadAll(chunkResp.Body)
	fmt.Printf("File content: %s\n", data)
}
