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
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type MutexInfo struct {
	mu *sync.Mutex
}

func NewMutexInfo(m *sync.Mutex) *MutexInfo {
	return &MutexInfo{
		mu: m,
	}
}

// Count returns the number of goroutines waiting on the mutex
func (m *MutexInfo) Count() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(m.mu)))
	v = v>>mutexWaiterShift + (v & mutexLocked)
	return int(v)
}

// IsLocked returns true if the mutex is locked
func (m *MutexInfo) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(m.mu)))
	return state&mutexLocked == mutexLocked
}

// IsWoken returns true if the mutex is woken
func (m *MutexInfo) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(m.mu)))
	return state&mutexWoken == mutexWoken
}

// IsStarving returns true if the mutex is starving
func (m *MutexInfo) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(m.mu)))
	return state&mutexStarving == mutexStarving
}
