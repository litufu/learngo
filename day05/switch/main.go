package main

import (
	"fmt"
	"time"
)

func main()  {
	i:=2
	switch i {
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("other")
	}
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's work day")
	}
	t := time.Now()
	switch {
	case t.Hour()>12:
		fmt.Println("it's morning")
	default:
		fmt.Println("it's afternoon")
	}
	whatAmI(1)
	whatAmI("nihao")
	whatAmI([]int{1,2,3})

}

func whatAmI(i interface{}){
	switch t:=i.(type) {
	case int:
		fmt.Println("i'm int")
	case string:
		fmt.Println("i'm string")
	default:
		fmt.Printf("i don't know,%T",t)


	}
}
