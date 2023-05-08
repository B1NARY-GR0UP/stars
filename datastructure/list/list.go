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

import "fmt"

// Node is a singly linked list
type Node struct {
	Val  int
	Next *Node
}

func New(val int) *Node {
	return &Node{
		Val:  val,
		Next: nil,
	}
}

func (node *Node) Append(val int) {
	head := node
	for head.Next != nil {
		head = head.Next
	}
	head.Next = New(val)
}

func (node *Node) Insert(val int) {
	newNode := New(val)
	newNode.Next = node.Next
	node.Next = newNode
}

func (node *Node) Remove(val int) {
	head := node
	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}
}

func (node *Node) String() string {
	var str string
	head := node
	for head != nil {
		str += fmt.Sprintf("%d -> ", head.Val)
		head = head.Next
	}
	str += "nil"
	return str
}

func IterSquareMutList(node *Node) {
	head := node
	for head != nil {
		head.Val *= head.Val
		head = head.Next
	}
}

func RecSquareMutList(node *Node) {
	// base case
	if node == nil {
		return
	}
	node.Val *= node.Val
	RecSquareMutList(node.Next)
}

// IterSquareList Square the elements
// Does not mutate the list
func IterSquareList(node *Node) *Node {
	if node == nil {
		return nil
	}
	res := New(node.Val * node.Val)
	curr := res
	node = node.Next
	for node != nil {
		curr.Next = New(node.Val * node.Val)
		curr = curr.Next
		node = node.Next
	}
	return res
}

// RecSquareList Square the elements
// Does not mutate the list
func RecSquareList(node *Node) *Node {
	if node == nil {
		return nil
	}
	res := New(node.Val * node.Val)
	res.Next = RecSquareList(node.Next)
	return res
}

func IterCatMutList(nodeA, nodeB *Node) *Node {
	if nodeA == nil {
		nodeA = nodeB
		return nodeA
	}
	if nodeB == nil {
		return nodeA
	}
	res := nodeA
	// 遍历完 res 指向最后一个 Node（不是 nil）
	for res.Next != nil {
		res = res.Next
	}
	for nodeB != nil {
		res.Next = nodeB
		res = res.Next
		nodeB = nodeB.Next
	}
	return nodeA
}

func RecCatMutList(nodeA, nodeB *Node) *Node {
	if nodeB == nil {
		return nodeA
	}
	if nodeA == nil {
		nodeA = nodeB
		return nodeA
	} else {
		nodeA.Next = RecCatMutList(nodeA.Next, nodeB)
		return nodeA
	}
}

func IterCatList(nodeA, nodeB *Node) *Node {
	res := new(Node)
	if nodeA == nil {
		res = nodeB
		return res
	}
	if nodeB == nil {
		res = nodeA
		return res
	}
	res = New(nodeA.Val)
	curr := res
	for nodeA != nil {
		curr.Next = New(nodeA.Val)
		nodeA = nodeA.Next
		curr = curr.Next
	}
	curr.Next = nodeB
	return res
}

func RecCatList(nodeA, nodeB *Node) *Node {
	if nodeA == nil && nodeB == nil {
		return nil
	}
	if nodeB == nil {
		return nodeA
	}
	if nodeA != nil {
		res := New(nodeA.Val)
		res.Next = RecCatList(nodeA.Next, nodeB)
		return res
	} else {
		res := New(nodeB.Val)
		res.Next = RecCatList(nodeA, nodeB.Next)
		return res
	}
}
