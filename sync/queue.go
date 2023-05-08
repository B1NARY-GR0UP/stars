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

package sync

import (
	"sync/atomic"
	"unsafe"
)

// Queue lock-free的queue
type Queue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

// 通过链表实现，这个数据结构代表链表中的节点
type node struct {
	value any
	next  unsafe.Pointer
}

func NewQueue() *Queue {
	n := unsafe.Pointer(&node{})
	return &Queue{head: n, tail: n}
}

// Enqueue 入队
func (q *Queue) Enqueue(v any) {
	n := &node{value: v}
	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		if tail == load(&q.tail) { // 尾还是尾
			if next == nil { // 还没有新数据入队
				if cas(&tail.next, next, n) { //增加到队尾
					cas(&q.tail, tail, n) //入队成功，移动尾巴指针
					return
				}
			} else { // 已有新数据加到队列后面，需要移动尾指针
				cas(&q.tail, tail, next)
			}
		}
	}
}

// Dequeue 出队，没有元素则返回nil
func (q *Queue) Dequeue() any {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head) { // head还是那个head
			if head == tail { // head和tail一样
				if next == nil { // 说明是空队列
					return nil
				}
				// 只是尾指针还没有调整，尝试调整它指向下一个
				cas(&q.tail, tail, next)
			} else {
				// 读取出队的数据
				v := next.value
				// 既然要出队了，头指针移动到下一个
				if cas(&q.head, head, next) {
					return v // Dequeue is done.  return
				}
			}
		}
	}
}

// 将unsafe.Pointer原子加载转换成node
func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// 封装CAS,避免直接将*node转换成unsafe.Pointer
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
