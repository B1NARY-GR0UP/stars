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

package mapreduce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := Map(items, func(a int) int {
		return a * 2
	})
	for i, item := range items {
		assert.Equal(t, item*2, res[i])
	}
}

func TestMapInSpace(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	MapInSpace(items, func(a int) int {
		return a * 2
	})
	for i, item := range items {
		assert.Equal(t, (i+1)*2, item)
	}
}

func TestReduce(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := Reduce(items, func(a, b int) int {
		return a + b
	})
	assert.Equal(t, 55, res)
}

func TestFilter(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := Filter(items, func(a int) bool {
		return a%2 == 0
	})
	assert.Equal(t, []int{2, 4, 6, 8, 10}, res)
}
