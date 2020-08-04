package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	data := "abc123!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	bs64,_ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(bs64))

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	us64,_:=base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(us64))
}