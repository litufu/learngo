package main

import "fmt"

type user struct {
	name string
	email string
}

func (u user)notify(){
	fmt.Printf("%s,%s\n",u.name,u.email)
}

func(u *user)changeEmail(email string){
	u.email = email
}

func main1()  {
	u1 := user{"lilei","lilei@email.com"}
	u2 := &user{"wangfei","wangfei@email.com"}
	u1.notify()
	u2.notify()
	u1.changeEmail("lilei@gewu.com")
	u2.changeEmail("wangfei@gewu.com")
	u1.notify()
	u2.notify()
}

