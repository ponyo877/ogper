package domain

import (
	"crypto/rand"
	"math/big"
)

var maxHexatridecimal7 = big.NewInt(78364164096)

type Hash struct {
	num int64
}

func NewHash() *Hash {
	n, err := rand.Int(rand.Reader, maxHexatridecimal7)
	if err != nil {
		panic(err)
	}
	return &Hash{n.Int64()}
}

// a-z0-9の36進数に変換
func (h *Hash) String() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 0)
	num := h.num
	for num > 0 {
		remainder := num % 36
		result = append([]byte{charset[remainder]}, result...)
		num /= 36
	}
	return string(result)
}
