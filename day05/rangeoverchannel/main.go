package main

import "fmt"

func main(){
	msg := make(chan string,2)
	msg <- "one"
	msg <- "two"
	close(msg)

	for item := range msg{
		fmt.Println(item)
	}
}

