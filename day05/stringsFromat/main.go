package main

import (
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func main(){
	p := point{1,2}
	fmt.Printf("%v\n",p)
	fmt.Printf("%T\n",p)
	fmt.Printf("%+v\n",p)
	fmt.Printf("%#v\n",p)
	fmt.Printf("%t\n",true)
	fmt.Printf("%d\n",123)
	fmt.Printf("%b\n",14)
	fmt.Printf("%c\n",33)
	fmt.Printf("%x\n",345)
	fmt.Printf("%f\n",78.9)
	fmt.Printf("%e\n",12340000.0)
	fmt.Printf("%E\n",12340000.0)
	fmt.Printf("%s\n","\"string\"")
	fmt.Printf("%q\n","\"string\"")
	fmt.Printf("%p\n",&p)
	fmt.Printf("|%6d|\n",123)
	fmt.Printf("|%6.2f|\n",45.9)
	fmt.Printf("|%-6.2f|\n",45.9)
	fmt.Printf("|%6s|\n","abc")
	fmt.Printf("|%-6s|\n","abc")
	s := fmt.Sprintf("a %s\n","string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr,"an %s\n","error")
}
