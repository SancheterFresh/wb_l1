package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mx sync.Mutex
	m  map[int]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int]int),
	}
}

func (sm *SafeMap) Write(k int, v int) {
	sm.mx.Lock()
	sm.m[k] = v
	fmt.Printf("Map key: %v = %v ", k, v)
	sm.mx.Unlock()
}

func Writer(sm *SafeMap, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := id; i < id+5; i++ {
		time.Sleep(2 * time.Second)
		sm.Write(i, i+10)

	}

}

func main() {
	safeMap := NewSafeMap()

	var wg sync.WaitGroup

	for i := 0; i <= 20; i += 5 {
		wg.Add(1)
		go Writer(safeMap, i, &wg)
	}

	wg.Wait()

}
