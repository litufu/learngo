package main

import (
	"fmt"
	"sort"
)

func main(){
	strings := []string{"asd","加速度","sakdjf","de3w"}
	sort.Strings(strings)
	fmt.Println(strings)

	ints := []int{3,5,2,3,7,2}
	sort.Ints(ints)
	fmt.Println(ints)
	fmt.Println(sort.IntsAreSorted(ints))
}
