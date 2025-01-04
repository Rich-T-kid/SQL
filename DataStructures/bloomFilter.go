package DataStructures

import (
	"hash"
	"hash/fnv"
	"math"
)

// BloomFilter represents a Bloom filter data structure
type BloomFilter struct {
	bitArray     []bool
	size         uint
	hashCount    uint
	hashFunc     hash.Hash64
	elementCount uint
}

// NewBloomFilter creates a new Bloom filter with the given parameters
func NewBloomFilter(expectedElements uint, falsePositiveRate float64) *BloomFilter {
	size := calculateSize(expectedElements, falsePositiveRate)
	hashCount := calculateHashCount(size, expectedElements)

	return &BloomFilter{
		bitArray:     make([]bool, size),
		size:         size,
		hashCount:    hashCount,
		hashFunc:     fnv.New64(),
		elementCount: 0,
	}
}

// calculateSize determines the optimal size of the bit array
func calculateSize(n uint, p float64) uint {
	m := -(float64(n) * math.Log(p)) / math.Pow(math.Log(2), 2)
	return uint(math.Ceil(m))
}

// calculateHashCount determines the optimal number of hash functions
func calculateHashCount(m, n uint) uint {
	k := (float64(m) / float64(n)) * math.Log(2)
	return uint(math.Ceil(k))
}

// getHashValues generates multiple hash values for an item
func (bf *BloomFilter) getHashValues(item []byte) []uint {
	bf.hashFunc.Reset()
	bf.hashFunc.Write(item)
	h64 := bf.hashFunc.Sum64()

	hashValues := make([]uint, bf.hashCount)
	for i := uint(0); i < bf.hashCount; i++ {
		// Use double hashing to generate multiple hash values
		hashValues[i] = uint((h64 + uint64(i)*uint64(h64>>32)) % uint64(bf.size))
	}
	return hashValues
}

// Add adds an item to the Bloom filter
func (bf *BloomFilter) Add(item []byte) {
	for _, pos := range bf.getHashValues(item) {
		bf.bitArray[pos] = true
	}
	bf.elementCount++
}

// Contains checks if an item might be in the set
func (bf *BloomFilter) Contains(item []byte) bool {
	for _, pos := range bf.getHashValues(item) {
		if !bf.bitArray[pos] {
			return false
		}
	}
	return true
}

// CurrentFalsePositiveRate calculates the current false positive rate
func (bf *BloomFilter) CurrentFalsePositiveRate() float64 {
	filledBits := 0
	for _, bit := range bf.bitArray {
		if bit {
			filledBits++
		}
	}

	fillRatio := float64(filledBits) / float64(bf.size)
	return math.Pow(1-math.Exp(-float64(bf.hashCount)*fillRatio), float64(bf.hashCount))
}

// ElementCount returns the number of elements added to the filter
func (bf *BloomFilter) ElementCount() uint {
	return bf.elementCount
}

// Size returns the size of the bit array
func (bf *BloomFilter) Size() uint {
	return bf.size
}

// HashCount returns the number of hash functions used
func (bf *BloomFilter) HashCount() uint {
	return bf.hashCount
}
