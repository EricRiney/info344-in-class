package main

import (
	"fmt"
	"os"
	"time"
)

const usage = `
usage:
	crawler <starting-url>
`

//worker is the worker function for our crawler. The main
//goroutine will start N of these running concurrently.
//The worker will read the next link out of the `linkq`,
//call getPageLinks to fetch and parse it, and then write
//the slice of links in that page to the `resultsq`
func worker(linkq chan string, resultsq chan []string) {
	//range over the linkq
	//as noted below, channel support the range
	//operator, returning the next item in the channel
	//(in this case, a link to fetch and parse)
	for link := range linkq {
		//call getPageLinks() to fetch the page, parse it,
		//and extract all of the hyperlinks in that page
		plinks, err := getPageLinks(link)
		//if we get an error, report it, but continue
		//so that the worker just gets the next item from
		//the `linkq`
		if err != nil {
			fmt.Printf("ERROR fetching %s: %v\n", link, err)
			continue
		}

		//write a message saying which link we fetched, and
		//how many links we found in that page
		fmt.Printf("%s (%d links)\n", link, len(plinks.Links))

		//sleep for half a second so that we don't pelt the
		//server with too many concurrent requests
		time.Sleep(time.Millisecond * 500)

		//if we found some links in the fetched page...
		if len(plinks.Links) > 0 {
			//use a new goroutine to write the links
			//to the resultq
			//we use a new goroutine here so that the
			//worker goroutine never gets blocked while
			//trying to write to the resultsq. Since the
			//main goroutine reads from that same queue
			//and writes each link to the `linkq`, which
			//the worker goroutine reads from, we could get
			//into a deadlock where each goroutine is waiting
			//for the other to write or read something.
			//using a new goroutine ensures the worker is
			//never blocked, so it can always read the next
			//link off of the linkq and thereby free-up the
			//main goroutine to read more off of the resultsq
			go func(links []string) {
				resultsq <- links
			}(plinks.Links)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	//number of worker go routines
	nWorkers := 1000

	//a channel that contains links to fetch and parse.
	//the second argument specifies how many links can
	//be written to the channel at a time.
	//if the channel is full, further writes to the channel
	//will block until a link is read from the channel
	linkq := make(chan string, 1000)

	//a channel for reporting results, which are all the
	//links in a page that was fetched and parsed.
	//the channel type here is []string, as we may find
	//multiple links in a given page
	resultsq := make(chan []string, 1000)

	//start our worker goroutines, passing
	//the linkq and the resultsq
	for i := 0; i < nWorkers; i++ {
		go worker(linkq, resultsq)
	}

	//write the URL the user supplied on the command
	//line to the linkq channel so that the first worker
	//can start crawling
	linkq <- os.Args[1]

	//construct a map to track links we've already seen
	//so that we don't fetch them again
	seen := map[string]bool{}

	//range over the resultsq
	//channels support the range operator, and will
	//return the next item in the channel for each
	//loop iteration
	for links := range resultsq {
		//since resultsq is a channel of []string
		//the `links` variable is a []string
		//so we can range over it like any slice
		for _, link := range links {
			//if we haven't seen this link before
			if !seen[link] {
				//mark that we have seen it
				seen[link] = true
				//and add it to the linkq,
				//which is the channel the workers
				//read from to get the next link
				//to fetch and parse
				linkq <- link
			}
		}
	}
}
