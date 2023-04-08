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

func isDigit(v int32) bool {
	// 0 - 9
	arr := []int32{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	left, right := 0, 9
	for left <= right {
		mid := (right-left)/2 + left
		num := arr[mid]
		if num == v {
			return true
		} else if num > v {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
