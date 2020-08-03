package main

import (
	"fmt"
	"time"
)

func main(){
	c1:=make(chan string ,1)
	go func() {
		time.Sleep(2*time.Second)
		c1<-"result1"
	}()

	select {
	case msg:=<-c1:
		fmt.Println(msg)
	case <-time.After(time.Second):
		fmt.Println("timeout1")
	}
	c2:=make(chan string ,1)
	go func() {
		time.Sleep(2*time.Second)
		c2<-"result2"
	}()

	select {
	case msg:=<-c2:
		fmt.Println(msg)
	case <-time.After(3*time.Second):
		fmt.Println("timeout2")
	}
}

