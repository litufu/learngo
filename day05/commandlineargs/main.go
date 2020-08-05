package main

import (
	"fmt"
	"os"
)

func main()  {
	all := os.Args
	args := all[1:]
	fmt.Println(len(all))
	arg := all[1]
	fmt.Println(all)
	fmt.Println(args)
	fmt.Printf("%T,%v",arg,arg)
}
