package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"net/url"

	"golang.org/x/net/html"
)

const headerContentType = "Content-Type"
const contentTypeHTML = "text/html"
const contentTypeJSON = "application/json; charset=utf-8"

//PageLinks contains summary information about a web page
type PageLinks struct {
	Title string   `json:"title"`
	Links []string `json:"links"`
}

//getPageLinks fetches PageSummary info for a given URL
func getPageLinks(URL string) (*PageLinks, error) {
	baseURL, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error resonse status code: %d", resp.StatusCode)
	}

	//if the requested URL is not an HTML page, just return a zero-value
	//PageLinks struct
	if !strings.HasPrefix(resp.Header.Get(headerContentType), contentTypeHTML) {
		return &PageLinks{}, nil
	}

	psum := &PageLinks{}
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		ttype := tokenizer.Next()
		if ttype == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				return psum, nil
			}
			return psum, tokenizer.Err()
		}

		//if this is a start tag token
		if ttype == html.StartTagToken {
			token := tokenizer.Token()
			//if this is the page title
			if token.Data == "title" {
				tokenizer.Next()
				psum.Title = tokenizer.Token().Data
			}

			//if this is a hyperlink
			if token.Data == "a" {
				//get the href attribute
				for _, attr := range token.Attr {
					//ignore bookmark links
					if attr.Key == "href" && !strings.HasPrefix(attr.Val, "#") {
						//parse the link, if error just continue
						link, err := url.Parse(attr.Val)
						if err != nil {
							continue
						}
						//if the link is not absolute (relative)
						//use the base URL's scheme and host
						if !link.IsAbs() {
							link.Scheme = baseURL.Scheme
							link.Host = baseURL.Host
						}
						psum.Links = append(psum.Links, link.String())
					}
				} //for all attributes
			} //if <a>
		} //if start tag
	} //for each token
} //getPageSummary()
