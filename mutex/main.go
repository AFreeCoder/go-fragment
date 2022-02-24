package main

import (
	"sync"
	"time"
)

func main() {
	// 死锁
	deadLock()
}

func deadLock() {
	var m1 sync.Mutex
	var m2 sync.Mutex
	var wg sync.WaitGroup

	var i = 0

	wg.Add(2)
	go func() {
		m1.Lock()
		defer m1.Unlock()

		i++
		time.Sleep(1 * time.Second)

		m2.Lock()
		defer m2.Unlock()
		wg.Done()
	}()

	go func() {
		m2.Lock()
		defer m2.Unlock()

		i++
		time.Sleep(1 * time.Second)

		m1.Lock()
		defer m1.Unlock()
		wg.Done()
	}()

	wg.Wait()
}
