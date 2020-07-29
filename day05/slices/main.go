package main

import "fmt"

func main()  {
	arr := make([]int,3)
	arr[0]=1
	arr[1]=2
	arr[2]=3
	fmt.Println(arr)
	arr = append(arr,4)
	fmt.Println(arr)
	fmt.Println(len(arr))

	s := make([]int,len(arr))
	copy(s,arr)
	fmt.Println(s)
	s[2]=7
	fmt.Println(s)
	fmt.Println(arr)

	d :=[]int{4,5,6}
	fmt.Println(d)

	e := make([][]int,3)
	for i:=0;i<len(e);i++{
		a:=make([]int,i+1)
		e[i]=a
		for j:=0;j<len(a);j++{
			a[j] = j+i
		}
	fmt.Println(e)
	}
}

