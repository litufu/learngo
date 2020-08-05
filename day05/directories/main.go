package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(err error){
	if err!=nil{
		panic(err)
    }
}

func main(){
	err := os.Mkdir("subdir",0755)
	check(err)
	defer os.Remove("subdir")

	createEmptyFile := func(name string){
		content := []byte{}
		err:= ioutil.WriteFile(name,content,0644)
		check(err)
	}

	createEmptyFile("subdir/file1")
	err =os.MkdirAll("subdir/parent/child",0755)
	check(err)
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c,err := ioutil.ReadDir("subdir/parent")
	check(err)

	fmt.Println("subdir/parent")
	for _,info := range c{
		fmt.Println(info.Name(),info.IsDir(),info.Mode())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)

	c,err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _,info := range c {
		fmt.Println(info.Name(),info.IsDir(),info.Mode())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)



}

func visit(p string,info os.FileInfo,err error) error{
	if err!=nil{
		return err
	}
	fmt.Println(p,info.Name(),info.IsDir(),info.Mode())
	return nil
}
