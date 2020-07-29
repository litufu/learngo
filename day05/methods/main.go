package main

import "fmt"

type reat struct {
	width,length int
}

func (r *reat) area()int{
	return r.width * r.length
}

func (r reat) perim()int{
	return 2*r.length+2*r.width
}
func main()  {
	r:=reat{3,4}
	fmt.Println(r.area())
	fmt.Println(r.perim())
	rp := r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
