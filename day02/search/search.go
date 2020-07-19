package search

import (
	"log"

	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string){
	feeds,err := RetrieveFeeds()
	if err != nil{
		log.Fatal(err)
	}

	resulsts := make(chan *Result)

	var waitGroup  sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _,feed := range feeds{
		matcher,exist := matchers[feed.Type]
		if ! exist{
			matcher = matchers["default"]
		}
		go func(feed *Feed,matcher Matcher) {
			Match(matcher,feed,searchTerm,resulsts)
			waitGroup.Done()
		}(feed,matcher)

	}
	go func() {
		waitGroup.Wait()

		close(resulsts)
	}()

	Display(resulsts)
}

func Register(feedType string,matcher Matcher){
	if _,exsists := matchers[feedType];exsists{
		log.Fatalln(feedType,"Matcher alrady register")
	}
	log.Println("register",feedType,"matcher")
	matchers[feedType] = matcher
}


