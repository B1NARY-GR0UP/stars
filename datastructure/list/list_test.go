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

package list

import (
	"fmt"
	"testing"
)

var testNodeA *Node
var testNodeB *Node

func init() {
	nodeA := New(1)
	nodeA.Append(2)
	nodeA.Append(3)
	nodeA.Append(4)
	nodeA.Append(5)
	testNodeA = nodeA

	nodeB := New(6)
	nodeB.Append(7)
	nodeB.Append(8)
	nodeB.Append(9)
	nodeB.Append(10)
	testNodeB = nodeB
}

func TestIterSquareMutList(t *testing.T) {
	IterSquareMutList(testNodeA)
	fmt.Println(testNodeA)
}

func TestRecSquareMutList(t *testing.T) {
	RecSquareMutList(testNodeA)
	fmt.Println(testNodeA)
}

func TestIterSquareList(t *testing.T) {
	node := IterSquareList(testNodeA)
	fmt.Println(node)
	fmt.Println(testNodeA)
}

func TestRecSquareList(t *testing.T) {
	node := RecSquareList(testNodeA)
	fmt.Println(node)
	fmt.Println(testNodeA)
}

func TestIterCatMutList(t *testing.T) {
	node := IterCatMutList(testNodeA, testNodeB)
	fmt.Println(node)
	fmt.Println(testNodeA)
	fmt.Println(testNodeB)
}

func TestRecCatMutList(t *testing.T) {
	node := RecCatMutList(testNodeA, testNodeB)
	fmt.Println(node)
	fmt.Println(testNodeA)
	fmt.Println(testNodeB)
}

func TestIterCatList(t *testing.T) {
	node := IterCatList(testNodeA, testNodeB)
	fmt.Println(node)
	fmt.Println(testNodeA)
	fmt.Println(testNodeB)
}

func TestRecCatList(t *testing.T) {
	node := RecCatList(testNodeA, testNodeB)
	fmt.Println(node)
	fmt.Println(testNodeA)
	fmt.Println(testNodeB)
}
