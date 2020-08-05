package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	os.Setenv("foo","abc")
	fmt.Println(os.Getenv("foo"))
	fmt.Println(os.Getenv("bar"))

	for _,e := range os.Environ(){
		pari := strings.SplitN(e,"=",2)
		fmt.Println(pari)
	}
}
