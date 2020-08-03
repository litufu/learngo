package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func main(){

	var readOp uint64
	var writeOp uint64

	type read struct {
		key int
		resp chan int
	}

	type write struct {
		key int
		value int
		resp chan bool
	}

	reads := make(chan read)
	writes := make(chan write)

	go func(){
		state:=make(map[int]int)
		for{
			select {
			case read:=<-reads:
				read.resp<-state[read.key]
			case write:=<-writes:
				state[write.key] = write.value
				write.resp<-true
			}
		}

	}()

	for i:=0;i<100;i++{
		go func(){
			for{
				read:=read{
					key:rand.Int(),
					resp:make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOp,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for j:=0;j<5;j++{
		go func(){
			for{
				write:=write{
					key:   rand.Int(),
					value: rand.Int(),
					resp:  make(chan bool),
				}
				writes<-write
				<-write.resp
				atomic.AddUint64(&writeOp,1)
				time.Sleep(time.Millisecond)
			}
		}()

	}

	time.Sleep(time.Second)

	fmt.Println(atomic.LoadUint64(&readOp))
	fmt.Println(atomic.LoadUint64(&writeOp))

}
