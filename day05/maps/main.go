package main

import "fmt"

func main()  {
	m := make(map[string]string)
	m["a"] = "a1"
	m["b"]= "b1"
	fmt.Println(m)
	fmt.Println(m["a"])
	fmt.Println(len(m))
	_,exist := m["c"]
	fmt.Println(exist)
	d :=map[string]int{"a":1,"b":2}
	fmt.Println(d)
	delete(d,"a")
	fmt.Println(d)
}
