package main

import (
	"crypto/md5"
	"fmt"
)

func main()  {
	str:="sha1 is a string"
	sh := md5.New()
	sh.Write([]byte(str))
	rs := sh.Sum([]byte("wo shi haoren"))
	fmt.Println(str)
	fmt.Printf("%x\n",rs)
}
