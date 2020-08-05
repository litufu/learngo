package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main()  {
	bi,lerr:= exec.LookPath("python")
	if lerr!=nil{
		panic(lerr)
	}

	args := []string{"python","-V"}

	env := os.Environ()

	exeErr := syscall.Exec(bi,args,env)
	if exeErr!=nil{
		panic(exeErr)
	}
}
