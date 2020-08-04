package main

import (
	"fmt"
	"time"
)

func main(){
	p := fmt.Println
	now := time.Now()
	p(now)

	then := time.Date(2017,1,2,1,2,1,11,time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Date())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	diff := then.Sub(now)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(now.Add(diff))
	p(now.Add(-diff))
}
