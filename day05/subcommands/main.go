package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	foocmd := flag.NewFlagSet("foo",flag.ExitOnError)
	barcmd := flag.NewFlagSet("bar",flag.ExitOnError)

	fooEnable := foocmd.Bool("enable",true,"fool enable bool")
	fooName := foocmd.String("name","","fool name string")

	barLevel := barcmd.Int("level",1,"bar level int")

	if len(os.Args)<2{
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		foocmd.Parse(os.Args[2:])
		fmt.Println(*fooEnable)
		fmt.Println(*fooName)
		fmt.Println(foocmd.Args())
	case "bar":
		barcmd.Parse(os.Args[2:])
		fmt.Println(*barLevel)
		fmt.Println(barcmd.Args())

	}
}
