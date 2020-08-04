package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error){
	if e!=nil{
		panic(e)
	}
}

func main(){
	content,err:=ioutil.ReadFile("D:/go/src/study/day05/readFiles/a.txt")
	check(err)
	fmt.Println(string(content))

	f,err := os.Open("D:/go/src/study/day05/readFiles/a.txt")
	check(err)

	b1 := make([]byte,5)
	n1,err := f.Read(b1)
	check(err)
	fmt.Println(string(b1[:n1]))

	o2,err := f.Seek(6,0)
	check(err)
	b2 := make([]byte,2)
	n2,err:=f.Read(b2)
	check(err)
	fmt.Println(o2,n2)
	fmt.Println(string(b2[:n2]))

	o3,err := f.Seek(6,0)
	check(err)
	b3:=make([]byte,3)
	n3,err:=io.ReadAtLeast(f,b3,3)
	check(err)
	fmt.Println(o3,n3)
	fmt.Println(string(b3))

	_,err = f.Seek(0,0)
	check(err)

	r4 := bufio.NewReader(f)
	b4,err:=r4.Peek(4)
	check(err)
	fmt.Println(string((b4)))

	defer f.Close()

}
