package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

const usage = `
usage:
	concur <data-dir-path> <search-string>
`

func processFile(filePath string, q string, ch chan []string) {
	//open the file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	//create a new scanner that reads from the file
	scanner := bufio.NewScanner(f)
	//create a new empty slice to hold the matches we find
	matches := []string{}
	for scanner.Scan() {
		//read each word
		//if it contains `q`
		//append it to `matches`
		word := scanner.Text()
		if strings.Contains(word, q) {
			matches = append(matches, word)
		}
	}

	//close the file
	f.Close()

	//write the matches slice to the channel
	//so that the caller can retrieve them
	ch <- matches
}

//processDir searches each file in the dirPath
//for the search string `q`
func processDir(dirPath string, q string) {
	//get all of the file names in the requested directory
	fileinfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	//make a channel that will receive a slices of
	//strings, one for each file we will process
	//these slices will be the matching words
	ch := make(chan []string, len(fileinfos))

	//for each of the files...
	for _, fi := range fileinfos {
		//call processFile on its own goroutine
		//to see the difference in performance, try
		//removing the `go` keyword, which will make
		//the processing happen serially instead of
		//concurrently
		go processFile(path.Join(dirPath, fi.Name()), q, ch)
	}

	//read from the channel once for each file we processed
	//and combine the results into one overall slice of strings
	totalMatches := []string{}

	//for each file we processed...
	for i := 0; i < len(fileinfos); i++ {
		//read the results from the channel
		matches := <-ch
		//append the matches to our overall slice
		//the `...` following the `matches` slice will
		//convert the slice elements into separate arguments
		//to the append() function
		totalMatches = append(totalMatches, matches...)
	}

	//sort the overall matches slice
	sort.Strings(totalMatches)

	//print all the matches to the terminal as a comma-delimeted list
	fmt.Println(strings.Join(totalMatches, ", "))
}

func main() {
	//this program will find all words that contain
	//a particular run of characters supplied by the user
	//the first argument is the directory containing the
	//data files, and the second argument is the run of
	//characters the user wants to search for
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	//the directory containing the data files
	dir := os.Args[1]

	//the query the user wants to search for
	//which will be a run of letters
	q := os.Args[2]

	fmt.Printf("processing directory %s...\n", dir)
	start := time.Now()
	processDir(dir, q)
	fmt.Printf("completed in %v\n", time.Since(start))
}
