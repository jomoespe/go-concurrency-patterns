package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Search func(query string) string

func fakeSearch(kind string) Search {
	return func(query string) string {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		return fmt.Sprintf("%s result for %q\n", kind, query)
	}
}

func Google(query string) (results []string) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

func main() {
}
