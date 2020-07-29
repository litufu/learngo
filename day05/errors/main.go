package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int,error){
	if arg==42{
		return 1,errors.New("can't be 42")
	}
	return arg,nil
}

type argError struct {
	arg int
	prob string
}

func (aE *argError) Error()string{
	return fmt.Sprintf("%d-%s",aE.arg,aE.prob)
}

func f2(arg int)(int,error){
	if arg==42{
		return 1,&argError{arg,"can't be 42!"}
	}
	return arg,nil
}

func main(){
	for _,i := range []int{7,42}{
		if r,e := f1(i); e!=nil{
			fmt.Println(e)
		}else{
			fmt.Println(r)
		}
		if r,e:=f2(i);e!=nil{
			fmt.Println(e)
		}else{
			fmt.Println(r)
		}
	}

	_,e := f2(42)
	if ae,ok := e.(*argError);ok{
		fmt.Println(ae.arg,ae.prob)
	}
}
