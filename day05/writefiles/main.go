package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error){
	if err!=nil{
		panic(err)
	}
}

func main(){
	content := []byte("wh't你很好")
	err := ioutil.WriteFile("./writefiles/a.txt",content,0644)
	check(err)

	f,err:=os.Create("./writefiles/b.txt")
	check(err)
	defer f.Close()

	n2,err := f.Write([]byte{115, 111, 109, 101, 10})
	check(err)
	fmt.Println(n2)

	n3,err := f.WriteString("wefcas的撒大V")
	check(err)
	fmt.Println(n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4,err := w.WriteString("abc")
	check(err)
	n5,err := w.Write([]byte("adskfj"))
	check(err)
	w.Flush()

	fmt.Println(n4,n5)




}
