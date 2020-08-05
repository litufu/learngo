package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main(){
	p := filepath.Join("d","ab","c")
	fmt.Println(p)
	fmt.Println(filepath.Join("dir//","filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	fmt.Println(filepath.Base(p))
	fmt.Println(filepath.Dir(p))

	fmt.Println(filepath.IsAbs("/d/d"))
	fmt.Println(filepath.IsAbs("d/d"))

	filename := "abc.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(filename,ext))

	rel,err:= filepath.Rel("a/b", "a/b/t/file")
	if err!=nil{
		panic(err)
	}
	fmt.Println(rel)

	rel1,err1:= filepath.Rel("a/b", "a/c/t/file")
	if err1!=nil{
		panic(err1)
	}
	fmt.Println(rel1)
}
