package main

import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(time.Second)
		c1 <-"c1"
	}()

	go func(){
		time.Sleep(2*time.Second)
		c2 <-"c2"
	}()
	for i:=0;i<2;i++{
		select {
		case c1:=<-c1:
			fmt.Println(c1)
		case c2:=<-c2:
			fmt.Println(c2)
		}
	}

}
