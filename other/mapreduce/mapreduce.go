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

// Map
// items: the items to map
// fn: the map function, fn(a) -> b, a is the item to map, b is the result
func Map[T any](items []T, fn func(T) T) []T {
	res := make([]T, len(items))
	for i, item := range items {
		res[i] = fn(item)
	}
	return res
}

// MapInSpace
// items: the items to map
// fn: the map function, fn(a) -> b, a is the item to map, b is the result
func MapInSpace[T any](items []T, fn func(T) T) {
	for i, item := range items {
		items[i] = fn(item)
	}
}

// Reduce
// items: the items to reduce
// fn: the reduce function, fn(a, b) -> c, a and b are the items to reduce, c is the result
func Reduce[T any](items []T, fn func(T, T) T) T {
	res := items[0]
	for i := 1; i < len(items); i++ {
		res = fn(res, items[i])
	}
	return res
}

// Filter
// items: the items to filter
// fn: the filter function, fn(a) -> bool
func Filter[T any](items []T, fn func(T) bool) []T {
	res := make([]T, 0)
	for _, item := range items {
		if fn(item) {
			res = append(res, item)
		}
	}
	return res
}
