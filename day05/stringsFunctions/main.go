package main

import (
	"fmt"
	s "strings"
	"unicode/utf8"
)

var p = fmt.Println

func main(){
	p("contains",s.Contains("test","st"))
	p("count",s.Count("test","st"))
	p("Index",s.Index("test","st"))
	p("HasPrefix",s.HasPrefix("test","te"))
	p("HasSuffix",s.HasSuffix("test","st"))
	p("Join",s.Join([]string{"a","b","c"},"-"))
	p("Repeat",s.Repeat("test",4))
	p("Replace1",s.Replace("test","st","ad",-1))
	p("Replace2",s.Replace("test","st","ad",1))
	p("Split",s.Split("t-e-s-t","-"))
	p("ToUpper",s.ToUpper("test"))
	p("ToLower",s.ToLower("TEST"))
	p()



	p(len("hello"))
	p(len("hello大家好"))
	p(utf8.RuneCountInString("hello大家好"))
	p("char","hello"[3])
	p("char","hello大家好"[5])
}
