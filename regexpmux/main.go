package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

const apiRoot = "/v1"

//RegExpMuxEntry represents an entry in the RegExpMux
type RegExpMuxEntry struct {
	pattern *regexp.Regexp
	handler http.Handler
}

//RegExpMux is a mux that matches requested resource
//paths using a regular expression
type RegExpMux struct {
	entries []*RegExpMuxEntry
}

//NewRegExpMux constructs and returns a new RegExpMux
func NewRegExpMux() *RegExpMux {
	return &RegExpMux{
		entries: []*RegExpMuxEntry{},
	}
}

//Handle adds a new HTTP handler to the mux, associated with the pattern
func (m *RegExpMux) Handle(pattern *regexp.Regexp, handler http.Handler) {
	m.entries = append(m.entries, &RegExpMuxEntry{
		pattern: pattern,
		handler: handler,
	})
}

//HandleFunc adds a handler function to the mux, associated with the pattern
func (m *RegExpMux) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	m.Handle(pattern, http.HandlerFunc(handler))
}

//ServeHTTP finds the appropriate handler given the requested URL.Path
//and calls that handler. If no match if found, it response with a
//not found 404 error
func (m *RegExpMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

//SpecificCourseHandler will handle requests for /v1/courses/course-id
func SpecificCourseHandler(w http.ResponseWriter, r *http.Request) {
}

//SpecificCourseRelationHandler will handle requests
//for /v1/courses/course-id/relation-type
func SpecificCourseRelationHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	addr := "localhost:4000"

	//create a new RegExpMux and use that
	//for the main server mux
	mux := NewRegExpMux()

	//TODO: add handlers

	fmt.Printf("listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
