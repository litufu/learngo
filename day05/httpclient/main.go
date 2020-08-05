package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main(){
	resp,err:=http.Get("http://www.baidu.com")
	if err!=nil{
		panic(err)
	}
	fmt.Println(resp.StatusCode)

	scanner := bufio.NewScanner(resp.Body)
	for i:=0;scanner.Scan()&&i<5;i++{
		fmt.Println(scanner.Text())
	}
	if err:=scanner.Err();err!=nil{
		panic(err)
	}
}
