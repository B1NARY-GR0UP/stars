package a

import (
	"time"
)

type ChanMutex struct {
	ch chan struct{}
}

func New() *ChanMutex {
	mu := &ChanMutex{
		ch: make(chan struct{}, 1),
	}
	mu.ch <- struct{}{}
	return mu
}

func (m *ChanMutex) Lock() {
	<-m.ch
}

// LockWithTimeout Lock and unlock in time t
func (m *ChanMutex) LockWithTimeout(t time.Duration) {
	timer := time.NewTimer(t)
	defer timer.Stop()
	for {
		select {
		case <-m.ch:
		case <-timer.C:
			m.Unlock()
			break
		}
	}
}

func (m *ChanMutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (m *ChanMutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

// TryLockWithTimeout Try to Lock in time t
func (m *ChanMutex) TryLockWithTimeout(t time.Duration) bool {
	timer := time.NewTimer(t)
	defer timer.Stop()
	select {
	case <-m.ch:
		return true
	case <-timer.C:
	}
	return false
}
