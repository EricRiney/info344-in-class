package main

import "math/rand"
import "fmt"
import "time"

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getRandomKey() string {
	return string(letters[rand.Intn(len(letters))])
}

func getter(c *Cache) {
	for {
		key := getRandomKey()
		fmt.Printf("getting key %s\n", key)
		c.Get(getRandomKey())
		time.Sleep(time.Millisecond * 500)
	}
}

func setter(c *Cache) {
	for {
		key := getRandomKey()
		fmt.Printf("setting key %s\n", key)
		c.Set(key, time.Now().String(), time.Second*2)
		time.Sleep(time.Millisecond * 600)
	}
}

func main() {
	//seed the random number generator
	rand.Seed(time.Now().Unix())

	//TODO: create a NewCache
	//and run a bunch of getter/setter
	//goroutines

	//A go program will exit when
	//the main() exits, and the
	//getter/setter functions are
	//running on their own goroutines,
	//so to keep the program running,
	//create a new unbuffered channel and
	//try to read from it;
	//since nothing is ever written to
	//the channel, this will block until
	//we kill the program with ctrl+c
	fmt.Println("hit ctrl+c to quit")
	quit := make(chan bool)
	<-quit
}
