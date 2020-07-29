package main

import (
	"fmt"
	"time"
)

func main(){
	channel := make(chan string)
	go func(){channel<-"hello"}()
	s := <-channel
	fmt.Println(s)

	messages := make(chan string,2)
	messages<-"hello"
	messages<-"world"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	pings := make(chan string,1)
	pongs := make(chan string,1)
	ping(pings,"nihao")
	pong(pings,pongs)
	fmt.Println(<-pongs)

	done := make(chan bool,1)

	go worker(done)
	<-done
}

func worker(done chan bool){
	fmt.Println("strating")
	time.Sleep(time.Second)
	fmt.Println("done")
	done<-true
}

func ping(pings chan<-string,msg string){
	pings<-msg
}

func pong(pings <-chan string,pongs chan<- string){
	msg:=<-pings
	pongs<-msg
}

