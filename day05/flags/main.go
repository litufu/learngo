package main

import (
	"flag"
	"fmt"
)

func main(){
	wordPtr := flag.String("word","foo","a string flag")
	numPtr := flag.Int("num",10,"a int flag")
	boolPtr := flag.Bool("fork",true,"a bool flag")

	var svar string
	flag.StringVar(&svar,"svar","bar","a svar flag")

	flag.Parse()

	fmt.Println(*wordPtr)
	fmt.Println(*numPtr)
	fmt.Println(*boolPtr)
	fmt.Println(svar)
	fmt.Println(flag.Args())
}
