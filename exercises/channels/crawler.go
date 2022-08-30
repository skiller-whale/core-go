package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

// Fetch URL `u` and return all the links found in <a href="..."> tags
func fetch(u string) (links []string) {
	pUrl, e := url.Parse(u)
	if e != nil {
		fmt.Println("😟🚮 Couldn't parse URL", u, " due to", e)
		return nil
	}
	resp, e := http.Get(u)
	if e != nil {
		fmt.Println("😟🔌 Couldn't fetch", u, "due to", e)
		return nil
	}

	for z := html.NewTokenizer(resp.Body); z.Next() != html.ErrorToken; {
		t := z.Token()
		if t.Type == html.StartTagToken && t.Data == "a" {
			for _, a := range t.Attr {
				if a.Key == "href" {
					u2, e := pUrl.Parse(a.Val)
					if e != nil {
						fmt.Println("😟 Couldn't parse relative url", a.Val)
					} else {
						links = append(links, u2.String())
					}
				}
			}
		}
	}
	return links
}

func UrlsToCrawlJobs(i []string, d int) (o []crawlJob) {
	for _, u := range i {
		o = append(o, crawlJob{Url: u, Depth: d})
	}
	return
}

// Read URLs to crawn on `s.in``, write the URLs we've found on `s.out`
func fetcher(s crawlerSupervision) {
	for j := range s.in {
		if j.Depth > 0 {
			s.out <- UrlsToCrawlJobs(fetch(j.Url), j.Depth-1)
		}
	}
}

type crawlJob struct {
	Url   string
	Depth int
}

type crawlerSupervision struct {
	in  chan crawlJob
	out chan []crawlJob
}

func main() {
	visited := make(map[string]bool)

	supervision := crawlerSupervision{
		in:  make(chan crawlJob),
		out: make(chan []crawlJob),
	}

	go fetcher(supervision)
	go fetcher(supervision)

	// Number of web requests "in flight"
	pending := 0
	// The jobs we've still got to do
	jobs := []crawlJob{{Depth: 3, Url: os.Args[1]}}

	// TODO: Write loop around select{} to complete the web fetch, using as many
	// simultaneous fetchers as we've started above 👆🏻
	//
	// It needs two cases:
	//
	// 1) To send nextJob to `supervision.in`
	//
	// 2) To receive results from `supervision.out`
	//
	// The receiver should:
	// * check whether the URLs returned have been visited, and ignore them if so
	// * check whether the returned jobs' depth is >0, add them back to the jobs list if so
	// * print out any URLs found
	//
	// Both sender and receiver should keep track of the number of pending requests
	// using `pending`, and store any unsent jobs in the `jobs` slice.
	//
}
