package main

import (
	"fmt"
	"hash"

	"github.com/spaolacci/murmur3"
	"github.com/google/uuid"
)

var mHash hash.Hash32

func init(){
	seed := 11
	mHash = murmur3.New32WithSeed(uint32(seed))
}

func murmurhash(key string, size int32) int32 {
    mHash.Reset()
    mHash.Write([]byte(key))
    result := mHash.Sum32() % uint32(size)
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

func (b *BloomFilter) Add(key string){
	idx := murmurhash(key, b.size)
	aIdx := idx/8
	bIdx := idx%8
	b.filter[aIdx] = b.filter[aIdx] | (1 << bIdx)
}

func (b *BloomFilter) Exists(key string) bool{
	idx := murmurhash(key, b.size)
	aIdx := idx/8
	bIdx := idx%8	
	return (b.filter[aIdx] & (1 << bIdx)) > 0
}

func main(){
	bloom := NewBloomFilter(100000)

	dataset := make([]string, 0)
	for i := 0; i<1000; i++{
		u := uuid.New()
		dataset = append(dataset, u.String())
	}

	for i:=0; i<500; i++{
		bloom.Add(dataset[i])
	}

	falsePositive := 0
	for i:=500; i<1000; i++{
		if bloom.Exists(dataset[i]){
			falsePositive++;
		}
	}
	fmt.Println(falsePositive)
	fmt.Println("False Positive Rate:", 100 * (float64(falsePositive)/ float64(len(dataset))))
}