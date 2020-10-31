package bloomfilter

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"math/big"
)

type bloomFilter struct {
	size   *big.Int
	bits   *big.Int
	hashes []hash.Hash
}

// BloomFilter interface (share an interface, not the implementation)
type BloomFilter interface {
	Add(n int)
	Has(b int) bool
}

func (bf *bloomFilter) findIndices(n int) (ret []int) {
	for _, hash := range bf.hashes {
		hash.Reset()
		hash.Write([]byte(fmt.Sprintln(n)))

		hv := new(big.Int)
		hv.SetBytes(hash.Sum(nil))

		hv.Mod(hv, bf.size)
		ret = append(ret, int(hv.Int64()))
	}
	return
}

// Add item to the bloom filter
func (bf *bloomFilter) Add(n int) {
	indices := bf.findIndices(n)
	for _, v := range indices {
		bf.bits.SetBit(bf.bits, v, 1)
	}
}

// Has test the item in bloom filter
func (bf *bloomFilter) Has(n int) bool {
	indices := bf.findIndices(n)
	count := 0
	for _, v := range indices {
		if bf.bits.Bit(v) == 1 {
			count++
		}
	}
	return count == len(indices)
}

// NewBloomFilter return a bloom filter
func NewBloomFilter(size int) BloomFilter {
	return &bloomFilter{
		big.NewInt(int64(size)),
		big.NewInt(0),
		[]hash.Hash{md5.New(), sha1.New(), sha256.New()},
	}
}
