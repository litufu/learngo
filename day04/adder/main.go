package main

import "fmt"

func adder() func() int{
	sum := 0
	return func() int{
		sum += 1
		return sum
	}
}

func main() {
	s:=adder()
	for i:=0;i<10;i++{
		fmt.Println(s())
	}
}
