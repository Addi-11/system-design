package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"sort"
)

var hashSpace int = 10000

func hashFunc(key string) int{
	hash := sha256.New()
	hash.Write([]byte(key)) // hash the key after convering it into bytes
	hashBytes := hash.Sum(nil)
	hashStr := hex.EncodeToString(hashBytes)
	hashInt := new(big.Int)
	hashInt.SetString(hashStr, 16)
	return int(hashInt.Mod(hashInt, big.NewInt(int64(hashSpace))).Int64())
}

func getServerConsistentHash(key string, backendServers []*OriginServer) *OriginServer{
	hash := hashFunc(key)

	// binary search to get idx >= hash
	idx := sort.Search(len(backendServers), func(i int) bool{return backendServers[i].pos >= hash})
	if idx == len(backendServers){
		idx = 0
	}

	return backendServers[idx]
}
