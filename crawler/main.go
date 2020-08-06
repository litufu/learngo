package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func check(err error){
	if err!=nil{
		panic(err)
	}
}

func main(){
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	check(err)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		fmt.Println("error get")
		return
	}

	charsetEncoding := determinCharset(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,charsetEncoding.NewDecoder())
	bytes, err := ioutil.ReadAll(utf8Reader)
	check(err)
	fmt.Println(string(bytes))

}

func determinCharset(reader io.Reader)encoding.Encoding{
	r := bufio.NewReader(reader)
	content,err := r.Peek(1024)
	check(err)
	encoding1, _, _ := charset.DetermineEncoding(content, "")
	return encoding1
}
