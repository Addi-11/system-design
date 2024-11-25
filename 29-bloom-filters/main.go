package main

import (
	"fmt"
	"hash"
	"math/rand"

	"github.com/spaolacci/murmur3"
	"github.com/google/uuid"
)

var hashfns[] hash.Hash32

func init(){
	for i := 0; i < 100; i++{
		hashfns = append(hashfns, murmur3.New32WithSeed(uint32(rand.Uint32())))
	}
}

func murmurhash(key string, size int32, hashIdx int) int32 {
    hashfns[hashIdx].Reset()
    hashfns[hashIdx].Write([]byte(key))
    result := hashfns[hashIdx].Sum32() % uint32(size)
    return int32(result)
}

type BloomFilter struct{
	filter []uint8
	size int32
}

func NewBloomFilter(size int32) *BloomFilter{
	return &BloomFilter{
		filter: make([]uint8, size),
		size: size,
	}
}

func (b *BloomFilter) Add(key string, numHashfns int){
	for i := 0; i < numHashfns; i++ {
		idx := murmurhash(key, b.size, i)
		aIdx := idx/8
		bIdx := idx%8
		b.filter[aIdx] = b.filter[aIdx] | (1 << bIdx)
	}
}

func (b *BloomFilter) Exists(key string, numHashfns int) bool{
	for i:= 0; i < numHashfns; i++{
		idx := murmurhash(key, b.size, i)
		aIdx := idx/8
		bIdx := idx%8	
		exsit := b.filter[aIdx] & (1 << bIdx) > 0
		if !exsit{
			return false
		}
	}
	return true
}

func main(){

	// different bloom filter sizes
	// sizes := []int32{16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144}

	dataset := make([]string, 0)
	for i := 0; i<5000; i++{
		u := uuid.New()
		dataset = append(dataset, u.String())
	}
	bloom := NewBloomFilter(30000) // constant size bloom filter

	for num := range(hashfns){
		
		j, diff := 2000, 1000
		// Adding values to the filters
		for i:=0; i<j; i++{
			bloom.Add(dataset[i], num)
		}

		// checking for values not in dataset
		falsePositive := 0
		for i:=j; i < j+diff; i++{
			if bloom.Exists(dataset[i], num){
				falsePositive++;
			}
		}
		// fmt.Printf("False Positive Rate: %f, Filter Size: %d\n", float64(falsePositive)/ float64(diff), size)
		fmt.Printf("False Positive Rate: %f, Num of hash Functions: %d\n", float64(falsePositive)/ float64(diff), num + 1)
	}
}