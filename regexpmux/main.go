package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

const apiRoot = "/v1"

var (
	//matches `/courses/unique-id`
	apiSpecificCourse = regexp.MustCompile(apiRoot + "/courses/([^/]+)$")
	//matches `/courses/unique-id/relation`
	apiSpecificCourseRelation = regexp.MustCompile(apiRoot + "/courses/([^/]+)/([^/]+)$")
)

//RegExpMuxEntry represents an entry in the RegExpMux
type RegExpMuxEntry struct {
	//a regular expression to compare against the requested path
	pattern *regexp.Regexp
	//a handler to call if the request path matches the pattern
	handler http.Handler
}

//RegExpMux is a mux that matches requested resource
//paths using a regular expression
type RegExpMux struct {
	//a slice of entries that have been added to the mux
	entries []*RegExpMuxEntry
}

//NewRegExpMux constructs and returns a new RegExpMux
func NewRegExpMux() *RegExpMux {
	return &RegExpMux{
		//initialize entries to a zero-length slice
		entries: []*RegExpMuxEntry{},
	}
}

//Handle adds a new HTTP handler to the mux, associated with the pattern
func (m *RegExpMux) Handle(pattern *regexp.Regexp, handler http.Handler) {
	//append a new entry, setting the pattern and handler
	m.entries = append(m.entries, &RegExpMuxEntry{
		pattern: pattern,
		handler: handler,
	})
}

//HandleFunc adds a handler function to the mux, associated with the pattern
func (m *RegExpMux) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	//call Handle() converting the handler function to an http.Handler
	m.Handle(pattern, http.HandlerFunc(handler))
}

//ServeHTTP finds the appropriate handler given the requested URL.Path
//and calls that handler. If no match if found, it response with a
//not found 404 error. This method makes our RegExpMux an http.Handler
func (m *RegExpMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handled := false

	//iterate over the entries...
	for _, entry := range m.entries {
		//if the entry's pattern matches the requested URL path...
		if entry.pattern.MatchString(r.URL.Path) {
			//call the handler
			entry.handler.ServeHTTP(w, r)
			handled = true
		}
	}

	//if no entry matched, respond with a 404
	if !handled {
		http.Error(w, "not found", http.StatusNotFound)
	}
}

//SpecificCourseHandler will handle requests for /v1/courses/course-id
func SpecificCourseHandler(w http.ResponseWriter, r *http.Request) {
	//to illustrate how you can use regular expressions to extract the
	//course ID portion of the path, we do the following...
	//find all of the capture-group matches
	matches := apiSpecificCourse.FindStringSubmatch(r.URL.Path)

	//the first match (index 0) will be the entire path
	//the second match (index 1) will be the course unique identifier
	msg := fmt.Sprintf("you asked for course %s", matches[1])
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(msg))
}

//SpecificCourseRelationHandler will handle requests
//for /v1/courses/course-id/relation-type
func SpecificCourseRelationHandler(w http.ResponseWriter, r *http.Request) {
	//to illustrate how you can use regular expressions to extract the
	//course ID and relation portion of the path, we do the following...
	//find all of the capture-group matches
	matches := apiSpecificCourseRelation.FindStringSubmatch(r.URL.Path)

	//the first match (index 0) will be the entire path
	//the second match (index 1) will be the course unique identifier
	//the third match (index 2) will be the relation type
	msg := fmt.Sprintf("you asked for the %s of course %s", matches[2], matches[1])
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(msg))
}

func main() {
	addr := "localhost:4000"

	//create a new RegExpMux and use that
	//for the main server mux
	mux := NewRegExpMux()

	//add handlers
	//first parameter is regular expression
	//second parameter is handler function
	mux.HandleFunc(apiSpecificCourse, SpecificCourseHandler)
	mux.HandleFunc(apiSpecificCourseRelation, SpecificCourseRelationHandler)

	fmt.Printf("listening at %s...\n", addr)
	//use our RegExpMux as the main server mux
	log.Fatal(http.ListenAndServe(addr, mux))
}
