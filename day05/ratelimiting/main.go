package main

import (
	"fmt"
	"time"
)

func main(){
	//每隔200毫秒处理异常请求
	requests := make(chan int,5)
	for i:=0;i<5;i++{
		requests<-i
	}
	close(requests)
	timeTicer1 := time.Tick(time.Millisecond*200)
	for req:= range requests{
		<-timeTicer1
		fmt.Println(req,time.Now())
	}

//	每隔200毫秒处理3个请求
	limiter := make(chan time.Time,3)

	newRequest := make(chan int,5)
	for m:=0;m<3;m++{
		limiter<-time.Now()
	}

	go func(){
		for t:=range time.Tick(time.Millisecond*500){
			limiter <- t
		}
	}()

	for j:=0;j<5;j++{
		newRequest<-j
	}
	close(newRequest)




	for j:= range newRequest{
		<-limiter
		fmt.Println("j",j,time.Now())
	}




}
