package main

import (
	"fmt"
	"time"
)

func main()  {
	 t:=time.Now()
	 fmt.Println(t.Format(time.RFC3339))

	 t,err:= time.Parse(time.RFC3339,"2020-08-04T14:47:40+08:00")
	 if err !=nil{
	 	fmt.Println(err)
	 }
	 fmt.Println(t)

	 fmt.Println(t.Format("3:04PM"))
	 form := "3 04 PM"
	 t2, e := time.Parse(form, "8 41 PM")
	 fmt.Println(e)
	 fmt.Println(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	fmt.Println(e)

}
