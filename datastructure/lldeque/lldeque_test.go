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

package lldeque

import (
	"fmt"
	"testing"
)

var testLL *LLDeque[int]

func init() {
	ll := New[int]()
	ll.AddFirst(1)
	ll.AddFirst(2)
	ll.AddFirst(3)
	ll.AddFirst(4)
	ll.AddFirst(5)
	testLL = ll
}

func TestNewWithCopy(t *testing.T) {
	ll := New[int]()
	ll.AddFirst(1)
	ll.AddFirst(2)
	ll.AddFirst(3)
	ll.AddFirst(4)
	ll.AddFirst(5)
	fmt.Println(ll)
	newLL := NewWithCopy[int](ll)
	fmt.Println(newLL)
}

func TestAddFirst(t *testing.T) {
	ll := New[int]()
	ll.AddFirst(1)
	ll.AddFirst(2)
	ll.AddFirst(3)
	fmt.Println(ll)
}

func TestAddLast(t *testing.T) {
	fmt.Println(testLL)
	testLL.AddFirst(-1)
	testLL.AddLast(-3)
	fmt.Println(testLL)
}

func TestDelFirst(t *testing.T) {
	fmt.Println(testLL)
	testLL.DelFirst()
	testLL.DelFirst()
	fmt.Println(testLL)
}

func TestDelLast(t *testing.T) {
	fmt.Println(testLL)
	testLL.DelLast()
	testLL.DelLast()
	fmt.Println(testLL)
}

func TestIterGet(t *testing.T) {
	fmt.Println(testLL)
	res := testLL.IterGet(4)
	fmt.Println(res)
}

func TestRecGet(t *testing.T) {
	fmt.Println(testLL)
	res := testLL.RecGet(3)
	fmt.Println(res)
}
