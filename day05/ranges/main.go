package main

import "fmt"

func main()  {
	sl := []int{1,2,3}
	for _,v := range sl{
		fmt.Println(v)
	}
	for k,v :=range sl{
		if v==3{
			fmt.Println(k)
		}
	}
	m :=map[string]string{"a":"a1","b":"b1"}
	for k,v := range m{
		fmt.Println(k,v)
	}
	for k := range m{
		fmt.Println(k)
	}
}
