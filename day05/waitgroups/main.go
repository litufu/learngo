package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("worker",id,"starting")
	time.Sleep(time.Second)
	fmt.Println("worker",id,"end")
}

func main()  {
	var wg sync.WaitGroup

	for i:=1;i<=5;i++{
		wg.Add(1)
		go worker(i,&wg)
	}

	wg.Wait()
}
