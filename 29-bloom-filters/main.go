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
	return b.filter[aIdx] & (1 << bIdx) > 0
}

func main(){

	sizes := []int32{16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144}

	dataset := make([]string, 0)
	for i := 0; i<5000; i++{
		u := uuid.New()
		dataset = append(dataset, u.String())
	}

	for _, size := range(sizes){
		bloom := NewBloomFilter(size)
		j, diff := 2000, 1000

		// Adding values to the filters
		for i:=0; i<j; i++{
			bloom.Add(dataset[i])
		}

		// checking for values not in dataset
		falsePositive := 0
		for i:=j; i < j+diff; i++{
			if bloom.Exists(dataset[i]){
				falsePositive++;
			}
		}
		fmt.Printf("False Positive Rate: %f, Filter Size: %d\n", float64(falsePositive)/ float64(diff), size)
	}
}