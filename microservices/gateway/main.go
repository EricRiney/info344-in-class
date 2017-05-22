package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

const defaultPort = "443"

func def(val, def string) string {
	if len(val) > 0 {
		return val
	}
	return def
}

//User represents a user in the system
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//getUser will return the currently authenticated
//User using information in the *http.Request
func getUser(r *http.Request) *User {
	//this is where you'd use your sessions
	//library to get the session state
	//and return the currently authenticated
	//user, but for purposes of this demo,
	//just return a test user
	return &User{
		ID:        "123456789",
		FirstName: "Test",
		LastName:  "User",
	}
}

//getServiceProxy returns a ReverseProxy for a microservice
//given the services address (host:port)
func getServiceProxy(svcAddr string) *httputil.ReverseProxy {
	instances := strings.Split(svcAddr, ",")
	nextInstance := 0

	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			user := getUser(r)

			//reset the scheme and Host of
			//the request URL
			r.URL.Scheme = "http"
			r.URL.Host = instances[nextInstance]
			nextInstance++
			nextInstance = nextInstance % len(instances)

			j, _ := json.Marshal(user)
			r.Header.Add("X-User", string(j))
		},
	}
}

func main() {
	port := def(os.Getenv("PORT"), defaultPort)
	host := os.Getenv("HOST")
	addr := fmt.Sprintf("%s:%s", host, port)
	certpath := def(os.Getenv("CERTPATH"), "./tls/fullchain.pem")
	keypath := def(os.Getenv("KEYPATH"), "./tls/privkey.pem")

	helloSvcAddr := os.Getenv("HELLOSVCADDR")
	if len(helloSvcAddr) == 0 {
		log.Fatal("you must supply a value for HELLOSVCADDR")
	}
	http.Handle("/hello", getServiceProxy(helloSvcAddr))

	fmt.Printf("gateway is listening at https://%s...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, certpath, keypath, nil))
}
