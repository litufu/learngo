package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	secs := now.Unix()
	nosecs := now.UnixNano()
	millis := nosecs /10000
	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(nosecs)
	fmt.Println(millis)

	fmt.Println(time.Unix(secs,0))
	fmt.Println(time.Unix(0,nosecs))
}
