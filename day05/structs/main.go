package main

import "fmt"

type Person struct {
	name string
	age int
}

func newPerson(name string) *Person{

	p:= Person{name,42}
	p.age=42
	return &p
}

func main()  {
	fmt.Println(Person{"bob",20})
	fmt.Println(Person{name:"lili",age:32})
	fmt.Println(Person{name:"lili"})
	fmt.Println(&Person{name:"lili",age:43})
	fmt.Println(newPerson("jon"))
	s := Person{"lili",34}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp)
}