package main

import (
	"fmt"
	"study/day04/retriver/mock"
	real2 "study/day04/retriver/real"
)

type Retriver interface {
	Get(url string) string
}

type Post interface {
	Post(url string,form map[string]string) string
}

type RetriverPoster interface {
	Retriver
	Post
}

func download(r Retriver) string{
	return r.Get("http://www.baidu.com")
}

func main(){
	var r Retriver
	r = &mock.Retriver{"this is a fake get"}

	var s RetriverPoster
	s = &mock.Retriver{Contents:"this is a new faker"}

	//r = &real2.Retriver{}
	fmt.Println(r.Get("http://www.baidu.com"))
	fmt.Println(s.Get("http://www.baidu.com"))
	s.Post("http://www.baidu.com", map[string]string{
		"contents":"another one",
	})
	fmt.Println(s.Get("http://www.baidu.com"))

	//realRetriver := r.(*real2.Retriver)
	//fmt.Printf("%T,%v",realRetriver,realRetriver)
	//if mockRetriver,ok := r.(mock.Retriver);ok{
	//	fmt.Printf("%T,%v",mockRetriver,mockRetriver)
	}


	//inspect(r)

//}

func inspect(r Retriver) {
	switch r.(type) {
	case *mock.Retriver:
		fmt.Printf("this is %T %V", r, r)
	case *real2.Retriver:
		fmt.Printf("this is %T %V", r, r)
	}
}
