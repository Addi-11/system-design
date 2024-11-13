package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"sort"
	"fmt"
	"log"
)

type StorageNode struct {
	pos int
	name string
	address string
}


var storageNodes []*StorageNode // nodes sorted by position in hash space 
var totalSlots int = 10000

func hash_fn(key string) int{
	h := sha256.New()
	h.Write([]byte(key)) // write bytes to the hash function
	hashBytes := h.Sum(nil)
	hashStr := hex.EncodeToString(hashBytes) // get hash in hexadecimal, and get the string

	hashInt := new(big.Int)
	hashInt.SetString(hashStr, 16)
	
	return int(hashInt.Mod(hashInt, big.NewInt(int64(totalSlots))).Int64())
}

func addNode(node *StorageNode){
	if len(storageNodes) >= totalSlots{
		log.Fatal("Hash map is full")
		return
	}
	node.pos = hash_fn(node.name)

	// binary search position to insert
	idx := sort.Search(len(storageNodes), func(i int) bool { return storageNodes[i].pos >= node.pos })
	// insert the node at the position
	storageNodes = append(storageNodes[:idx], append([]*StorageNode{node}, storageNodes[idx:]...)...)
	fmt.Printf("Added node %s at pos %d\n", node.name, node.pos)
}

func removeNode(nodeName string) {
	// binary search to find the node by its hash position
	index := sort.Search(len(storageNodes), func(i int) bool { return storageNodes[i].name == nodeName })

	if index < len(storageNodes) && storageNodes[index].name == nodeName {
		storageNodes = append(storageNodes[:index], storageNodes[index+1:]...)
		fmt.Printf("Node %s removed\n", nodeName)
	} else {
		fmt.Printf("Node %s not found\n", nodeName)
	}
}

func findNode(text string) *StorageNode{
	pos := hash_fn(text)

	// binary search idx in storageNodes, idx >= key
	idx :=  sort.Search(len(storageNodes), func(i int) bool { return storageNodes[i].pos >= pos })
	if idx == len(storageNodes){
		idx = 0
	}
	return storageNodes[idx]
}

func viewStorageNodes(){
	fmt.Println("\n------------------------------")
	fmt.Println("STORAGE NODE STATS")
	for _, n := range(storageNodes){
		fmt.Printf("Node %s is placed at pos %d\n", n.name, n.pos)
	}
	
	fmt.Println("------------------------------\n")
}

func main(){
	nodes := []*StorageNode{
		&StorageNode{name: "A", address: "8001"},
		&StorageNode{name: "B", address: "8002"},
		&StorageNode{name: "C", address: "8003"},
		&StorageNode{name: "D", address: "8004"},
	}

	for _, n := range(nodes){
		addNode(n)
	}

	viewStorageNodes()

	text := "testKey"
	node := findNode(text)
	fmt.Printf("Node for key (%s): %s, pos %d\n", text, node.name, node.pos)

	removeNode("D")

	viewStorageNodes()

	// see where the key is inserted now
	node = findNode(text)
	fmt.Printf("Node for key (%s): %s, pos %d\n", text, node.name, node.pos)
}