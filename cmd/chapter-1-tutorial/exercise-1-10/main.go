/*
	Find a website that produces a large amount of data. Investigate caching by
	running fetchall twice in succession to see whether the reported time changes much. Do
	you get the same content each time? Modify fetchall to print its output to a file so it can be
	examined.
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)

	for idx, url := range os.Args[1:] {
		go fetch(url, idx, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, page int, ch chan<- string) {
	start := time.Now()

	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	f, err := os.Create("../../Projects/go-exercises/cmd/chapter-1-tutorial/exercise-1-10/page" + strconv.Itoa(page) + ".txt")

	if err != nil {
		ch <- fmt.Sprint(err)
	}

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	_, err = io.Copy(f, resp.Body)

	f.Close()

	resp.Body.Close() // don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %s", secs, url)
}
