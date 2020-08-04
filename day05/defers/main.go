package main

import (
	"fmt"
	"os"
)

func main(){
	f := createFIle("./defers/a.txt")
	defer closeFile(f)
	writeFile(f)

}

func createFIle(p string) *os.File{
	fmt.Println("creating file")
	f,err:=os.Create(p)
	if err!=nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("writing data")
	fmt.Fprintf(f,"data")
}

func closeFile(f *os.File){
	fmt.Println("closing file")
	err:= f.Close()
	if err!=nil{
		fmt.Fprintf(os.Stderr,"%v\n",err)
		os.Exit(1)
	}
}
