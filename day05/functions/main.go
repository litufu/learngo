package main

import "fmt"

func plus(a int,b int)int  {
	return a + b
}

func plusA(a,b,c int)int  {
	return a + b +c
}

func num() (int,int){
	return 3,5
}

func interSeq()func()int{
	sum :=0
	return func()int{
		sum +=1
		return sum
	}
}

func sum(nums ...int ){
	fmt.Println(nums)
	sum :=0
	for _,v :=range nums{
		sum +=v
	}
	fmt.Println(sum)
}

func fact(i int)  int{
	if i<=1{
		return 1
	}
	return i * fact(i-1)
}

func main(){
	s1 := plus(3,4)
	fmt.Println(s1)
	s2 := plusA(3,4,5)
	fmt.Println(s2)
	a,b:=num()
	fmt.Println(a,b)
	sum(1,2)
	sum(1,2,3)
	sl := []int{4,5,6,7}
	sum(sl...)

	nextInt:=interSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := interSeq()
	fmt.Println(nextInts())

	fmt.Println(fact(7))
}
