package main

import "fmt"

func main(){
	messages := make(chan string)
	signals:=make(chan string)

	select {
	case msg:=<-messages:
		fmt.Println(msg)
	default:
		fmt.Println("no msg")
	}

	msg := "abc"
	select {
	case messages<-msg:
		fmt.Println("send msg")
	default:
		fmt.Println("no send msg")
	}

	select {
	case msg:=<-messages:
		fmt.Println(msg)
	case sig:=<-signals:
		fmt.Println(sig)
	default:
		fmt.Println("no msg")
	}
}
