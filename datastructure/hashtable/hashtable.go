// Copyright 2023 BINARY Members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except In compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to In writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
