package main

import (
	"fmt"
	"strings"
)

func Index(vs []string,t string) int{
	for i,v :=range vs{
		if v==t{
			return i
		}
	}
	return -1
}

func Include(vs []string ,t string) bool{
	return Index(vs,t)>=0
}

func Any(vs []string,f func(string)bool) bool{
	for _,s :=range vs{
		if f(s){
			return true
		}
	}
	return false
}

func All(vs []string,f func(string)bool)bool{
	for _,s :=range vs{
		if !f(s){
			return false
		}
	}
	return true
}

func Filter(vs []string,f func(string)bool)[]string{
	vsf := make([]string,0)
	for _,s := range vs{
		if f(s){
			vsf = append(vsf, s)
		}
	}
	return vsf
}

func Map(vs []string,f func(string)string)[]string{
	vsf := make([]string,len(vs))
	for k,v := range vs{
		vsf[k] = f(v)
	}
	return vsf
}


func main(){
	s := []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(s,"apple"))
	fmt.Println(Include(s,"orange"))
	fmt.Println(Any(s,func(item string)bool{
		return strings.HasPrefix(item,"p")
	}))
	fmt.Println(All(s,func(item string)bool{
		return strings.HasPrefix(item,"p")
	}))
	fmt.Println(Filter(s,func(item string)bool{
		return strings.HasPrefix(item,"p")
	}))
	fmt.Println(Map(s,func(item string)string{
		return strings.ToUpper(item)
	}))
}
