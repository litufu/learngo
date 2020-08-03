package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main(){
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOps int32
	var writeOps uint64

	for i:=0;i<100;i++{
		go func(){
			for{
				key := rand.Int()
				mutex.Lock()
				_ = state[key]
				mutex.Unlock()
				atomic.AddInt32(&readOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadInt32(&readOps)
	fmt.Println(readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println(writeOpsFinal)

	mutex.Lock()
	fmt.Println(state)
	mutex.Unlock()
}
