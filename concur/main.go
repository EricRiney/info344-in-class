package main

import (
	"fmt"
	"os"
	"time"
)

const usage = `
usage:
	concur <data-dir-path>
`

func processFile(filePath string, ch chan int) {
	//TODO: open the file, scan each line,
	//do something with the word, and write
	//the results to the channel
}

func processDir(dirPath string) {
	//TODO: iterate over the files in the directory
	//and process each, first in a serial manner,
	//and then in a concurrent manner
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	dir := os.Args[1]

	fmt.Printf("processing directory %s...\n", dir)
	start := time.Now()
	processDir(dir)
	fmt.Printf("completed in %v\n", time.Since(start))
}
