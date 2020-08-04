package main

import (
	"fmt"
	"strconv"
)

func main(){
	i,_ := strconv.ParseInt("123",0,64)
	fmt.Println(i)
	f,_ := strconv.ParseFloat("123.45",64)
	fmt.Println(f)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
