package main

import (
	"fmt"
	"time"
)

func main(){
	timer1 := time.NewTimer(2*time.Second)

	<-timer1.C
	fmt.Println("timer1 is fire")

	time2 := time.NewTimer(time.Second)
	go func(){
		<-time2.C
		fmt.Println("time2 is fire")
	}()
	stop := time2.Stop()
	if stop{
		fmt.Println("time2 is stop")
	}
	time.Sleep(2*time.Second)
}
