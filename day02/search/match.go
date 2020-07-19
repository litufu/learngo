package search

import "log"

type Result struct {
	Field string
	Content string
}

type Matcher interface {
	Search(feed *Feed,searchTerm string)([]*Result,error)
}

func Match(matcher Matcher,feed *Feed,searchTerm string,results chan<-*Result){
	searchResults,err := matcher.Search(feed ,searchTerm)
	if err!=nil{
		log.Println(err)
		return
	}
	for _,result := range searchResults{
		results <- result
	}
}

// Display writes results to the console window as they

// are received by the individual goroutines.

func Display(results chan *Result) {

	// The channel blocks until a result is written to the channel.

	// Once the channel is closed the for loop terminates.

	for result := range results {

		log.Printf("%s:\n%s\n\n", result.Field, result.Content)

	}

}