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
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type ReentrantLock struct {
	sync.Mutex
	owner int64
	recur int64
}

func (l *ReentrantLock) Lock() {
	currID := l.getGoroutineID()
	if atomic.LoadInt64(&l.owner) == currID {
		l.recur++
		return
	}
	l.Mutex.Lock()
	atomic.StoreInt64(&l.owner, currID)
	l.recur = 1
}

func (l *ReentrantLock) Unlock() {
	currID := l.getGoroutineID()
	if atomic.LoadInt64(&l.owner) != currID {
		panic(fmt.Sprintf("unlock failed, you are %v, owner is %v", currID, l.owner))
	}
	l.recur--
	if l.recur != 0 {
		return
	}
	atomic.StoreInt64(&l.owner, -1)
	l.Mutex.Unlock()
}

func (l *ReentrantLock) getGoroutineID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.ParseInt(idField, 10, 0)
	if err != nil {
		panic(fmt.Sprintf("get goroutine id err: %v", err))
	}
	return id
}
