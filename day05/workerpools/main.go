package main

import (
	"fmt"
	"time"
)

func worker(id int,jobs <-chan int,results chan<- int){
	for job :=  range jobs{
		fmt.Println("worker",id,"starting work",job)
		time.Sleep(time.Second)
		fmt.Println("worker",id,"end work",job)
		results <- 2*job
	}

}

func main(){
	const numJobs = 5
	jobs := make(chan int,numJobs)
	results := make(chan int,numJobs)
	for i:=0;i<3;i++{
		go worker(i,jobs,results)
	}

	for j:=0;j<numJobs;j++{
		jobs<-j
	}
	close(jobs)

	for s:=0;s<numJobs;s++{
		<-results
	}
}
