package main

import "os"

func main(){
	//panic("a problem")

	_,err:=os.Create("/temp/a.txt")
	if err!=nil{
		panic(err)
	}
}
