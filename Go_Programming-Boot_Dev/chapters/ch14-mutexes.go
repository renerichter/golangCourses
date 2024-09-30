package main

import (
	"fmt"
	"sync"
)

// type safeCounter struct {
// 	counts map[string]int //map with type(key)=string and type(val)=int
// 	mux    *sync.Mutex    //pointer that is shared accross all goroutines
// }

// func (sc safeCounter) inc(key string) {
// 	sc.mux.Lock()
// 	defer sc.mux.Unlock()
// 	sc.slowIncrement(key) // called at different parts of execution code

// }
// func (sc safeCounter) val(key string) int {
// 	sc.mux.Lock()
// 	defer sc.mux.Unlock()
// 	return sc.counts(key)
// }

/* type safeCounter struct {
	counts map[string]int
	mux    *sync.Mutex
}

func (sc *safeCounter) slowIncrement(key string) {
	sc.counts[key]++
	time.Sleep(time.Millisecond * 10)
}

func (sc *safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc *safeCounter) val(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.counts[key]
} */
