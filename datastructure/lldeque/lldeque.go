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

import "fmt"

// LLDeque is a linked list implementation of a deque
type LLDeque[T any] struct {
	sentinel *Node[T]
	size     int
}

type Node[T any] struct {
	Prev, Next *Node[T]
	Item       T
}

func New[T any]() *LLDeque[T] {
	ll := &LLDeque[T]{
		sentinel: &Node[T]{},
	}
	ll.sentinel.Prev = ll.sentinel
	ll.sentinel.Next = ll.sentinel
	return ll
}

func NewWithCopy[T any](ll *LLDeque[T]) *LLDeque[T] {
	newLL := New[T]()
	for i := 0; i < ll.size; i++ {
		newLL.AddLast(ll.IterGet(i))
	}
	return newLL
}

func (ll *LLDeque[T]) Size() int {
	return ll.size
}

func (ll *LLDeque[T]) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LLDeque[T]) AddFirst(item T) {
	newNode := &Node[T]{
		Prev: ll.sentinel,
		Next: ll.sentinel.Next,
		Item: item,
	}
	ll.sentinel.Next.Prev = newNode
	ll.sentinel.Next = newNode
	ll.size++
}

func (ll *LLDeque[T]) AddLast(item T) {
	newNode := &Node[T]{
		Prev: ll.sentinel.Prev,
		Next: ll.sentinel,
		Item: item,
	}
	ll.sentinel.Prev.Next = newNode
	ll.sentinel.Prev = newNode
	ll.size++
}

func (ll *LLDeque[T]) DelFirst() T {
	var zero T
	if ll.IsEmpty() {
		return zero
	}
	delNode := ll.sentinel.Next
	delNode.Prev.Next = delNode.Next
	delNode.Next.Prev = delNode.Prev
	ll.size--
	return delNode.Item
}

func (ll *LLDeque[T]) DelLast() T {
	var zero T
	if ll.IsEmpty() {
		return zero
	}
	delNode := ll.sentinel.Prev
	delNode.Prev.Next = delNode.Next
	delNode.Next.Prev = delNode.Prev
	ll.size--
	return delNode.Item
}

func (ll *LLDeque[T]) IterGet(index int) T {
	var zero T
	if index < 0 || index >= ll.size {
		return zero
	}
	head := ll.sentinel.Next
	for i := 0; i < index; i++ {
		head = head.Next
	}
	return head.Item
}

func (ll *LLDeque[T]) RecGet(index int) T {
	var zero T
	if index < 0 || index >= ll.size {
		return zero
	}
	return ll.recGetHelper(ll.sentinel.Next, index)
}

func (ll *LLDeque[T]) recGetHelper(head *Node[T], index int) T {
	if index == 0 {
		return head.Item
	}
	return ll.recGetHelper(head.Next, index-1)
}

func (ll *LLDeque[T]) String() string {
	var str string
	tmp := ll.sentinel.Next
	for tmp != ll.sentinel {
		str += fmt.Sprintf("%v <=> ", tmp.Item)
		tmp = tmp.Next
	}
	return str
}
