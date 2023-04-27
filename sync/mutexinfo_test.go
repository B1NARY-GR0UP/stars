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
	"sync"
	"testing"
	"time"
)

func TestMutexInfo(t *testing.T) {
	var m sync.Mutex
	mi := NewMutexInfo(&m)
	for i := 0; i < 1000; i++ { // 启动1000个goroutine
		go func() {
			mi.mu.Lock()
			time.Sleep(time.Second)
			mi.mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t,  starving: %t\n", mi.Count(), mi.IsLocked(), mi.IsWoken(), mi.IsStarving())
}
