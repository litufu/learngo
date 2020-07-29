package main

import "fmt"

type react struct {
	height,length float64
}

type circle struct {
	radis float64
}

type geometry interface {
	area() float64
	perim()float64
}

func (r react)area()float64{
	return r.height * r.length
}

func (r react)perim()float64  {
	return 2*r.length+2*r.height
}

func (c circle)area()float64{
	return 3.14*c.radis*c.radis
}

func (c circle)perim()float64{
	return 2*3.14*c.radis
}

func measure(g geometry){
	fmt.Println(g.area())
	fmt.Println(g.perim())
	fmt.Println(g)
}

func main()  {
	r := react{3,4}
	c := circle{2}
	measure(r)
	measure(c)
}
