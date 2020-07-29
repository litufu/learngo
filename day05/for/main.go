package main

import "fmt"

func main()  {
	i := 1
	for i<3{
		i++
		fmt.Println(i)
	}
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	for {
		fmt.Println("hello")
		break
	}

	for i:=0;i<20;i++{
		if i%2==0{
			continue
		}
		fmt.Println(i)
	}
}
