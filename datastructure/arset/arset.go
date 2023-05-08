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

package arset

type ARSet[T comparable] struct {
	items []T
	size  int
}

func New[T comparable]() *ARSet[T] {
	return &ARSet[T]{
		items: make([]T, 0),
		size:  0,
	}
}

func (s *ARSet[T]) Size() int {
	return s.size
}

func (s *ARSet[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *ARSet[T]) Add(item T) {
	if !s.Contains(item) {
		s.items = append(s.items, item)
		s.size++
	}
}

func (s *ARSet[T]) Contains(item T) bool {
	for _, it := range s.items {
		if it == item {
			return true
		}
	}
	return false
}
