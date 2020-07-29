package main

import "fmt"

func main()  {
	var a [5]int
	a[4] = 0
	fmt.Println(a)
	var b [3][4]int
	for i:=0;i<3;i++{
		for j:=0;j<4;j++{
			b[i][j]=i*j
		}
	}
	fmt.Println(b)
	c :=[5]string{"a","b","c","d","e"}
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(c[3])
	fmt.Println(c[3:])
	fmt.Println(c[:3])
}
