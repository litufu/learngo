package main

import "fmt"

func val(value int){
	value = 0
}

func valptr(value *int){
	*value = 1
}

func main()  {
	a := 3
	fmt.Println(a)

	val(a)
	fmt.Println(a)
	valptr(&a)
	fmt.Println(a)

}

