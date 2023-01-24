package hashtable

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type (
	core     [][]string
	hashFunc func(s string) uint
)

// HashTable core struct
type HashTable struct {
	core     core
	hashFunc hashFunc
	initSize int // TODO: use make
	k        int // TODO: load factor
}

// NewHashTable TODO: use optional function style
func NewHashTable(initSize, k int) *HashTable {
	return &HashTable{}
}

func hash(key string) uint {
	h := md5.New()
	h.Write([]byte(key))
	hashString := hex.EncodeToString(h.Sum(nil)) // length = 32
	sb := strings.Builder{}
	for _, v := range hashString {
		if isDigit(v) {
			_, _ = fmt.Fprint(&sb, string(v))
		}
	}
	i, _ := strconv.Atoi(sb.String()[:3]) // may cause err
	return uint(i)
}
