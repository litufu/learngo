package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter,req *http.Request){
	fmt.Fprint(w,"hello\n")
}

func headers(w http.ResponseWriter,req *http.Request){
	for name,headers := range req.Header{
		for _,value:=range headers{
			fmt.Fprintf(w,"%v %v\n",name,value)
		}
	}
}

func main(){
	http.HandleFunc("/hello",hello)
	http.HandleFunc("/headers",headers)

	http.ListenAndServe(":8090",nil)
}
