package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c cache, ch chan string) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	defer close(ch)
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	// url already visited
	if _, ok := c.Get(url); ok {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	c.Set(url, true)

	if err != nil {
		ch <- err.Error()
		return
	}
	ch <- fmt.Sprintf("found: %s %q", url, body)
	results := make([]chan string, len(urls))
	for i, u := range urls {
		results[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, c, results[i])
	}

	for _, result := range results {
		for url := range result {
			ch <- url
		}
	}
	return
}

func main() {
	c := cache{
		visited: map[string]bool{},
		mux:     &sync.RWMutex{},
	}
	ch := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, c, ch)
	for url := range ch {
		fmt.Println(url)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

type safeMap[T any, K any] interface {
	Set(T, K)
	Get(T) (K, bool)
}

type cache struct {
	visited map[string]bool
	mux     *sync.RWMutex
}

func (s *cache) Set(url string, visited bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.visited[url] = visited
}

func (s *cache) Get(url string) (bool, bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	visited, ok := s.visited[url]
	return visited, ok
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
