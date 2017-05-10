package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//MessageEvent represents an event with a message
//and a timestamp
type MessageEvent struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

//TriggerEvent triggers a new MessageEvent. This is just a handy
//way to create new events for demo purposes. In a real app, you
//would create and broacast events in response to various handler
//actions, e.g., new user sign-up, post of a new message, etc.
func TriggerEvent(w http.ResponseWriter, r *http.Request) {
	//CORS headers to allow cross-origin requests
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Request-Method", "POST")
	w.Header().Add("Access-Control-Request-Headers", "Content-Type")

	//TODO: create a new MessageEvent with a hard-coded message
	//and the current time for CreatedAt
	//Then pass the MessageEvent to the `.Notify()` method of your notifier
	//so that the event gets broadcasted to all web socket clients
}

//WebSocketUpgradeHandler handles websocket upgrade requests
func WebSocketUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: upgrade this request to a web socket connection
	//see https://godoc.org/github.com/gorilla/websocket#hdr-Overview
	//NOTE that by default, the websocket package will reject
	//cross-origin upgrade requests, so make sure you set the
	//CheckOrigin field of the Upgrader to allow upgrades from
	//any origin.
	//See https://godoc.org/github.com/gorilla/websocket#hdr-Origin_Considerations

	//after upgrading, use the `.AddClient()` to add the new
	//connection to your notifier

}

func main() {
	addr := "localhost:4000"

	//TODO: create a NewNotifier and call
	//its .Start() method on a new goroutine
	// go mynotifier.Start()

	//your handlers will need access to this notifer
	//instance, so share it with them somehow

	http.HandleFunc("/v1/ws", WebSocketUpgradeHandler)
	http.HandleFunc("/v1/trigger", TriggerEvent)

	fmt.Printf("listening at %s...\n", addr)
	fmt.Printf("test the server by opening the websockclient/index.html page\n")
	fmt.Printf("in a few different browser tabs\n")
	log.Fatal(http.ListenAndServe(addr, nil))
}
