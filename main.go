package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func filename() string {
	return fmt.Sprintf("%s_%d", time.Now().Month(), time.Now().Day())
}

func filepath() string {
	return fmt.Sprintf("events/%s", filename())
}

func cachedFile(path string) bool {
	fileExists := false
	if _, err := os.Stat(path); err == nil {
		fileExists = true
	}
	return fileExists
}

func downloadFile() {
	os.Mkdir("events", 0777)
	file, err := os.Create(filepath())

	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", filename())
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to get URL:\n%s", err)
		os.Exit(1)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		println(err)
	}
	defer file.Close()
	defer response.Body.Close()
}

func readCache() io.ReadCloser {
	file, err := os.Open(filepath())
	if err != nil {
		fmt.Printf("Failed to open file:\n%s", err)
		os.Exit(1)
	}
	return file
}

func eventsHtml() io.ReadCloser {
	if cachedFile(filepath()) == false {
		downloadFile()
	}
	return readCache()
}

func events() []string {
	var events []string
	var text []byte
	body := eventsHtml()
	z := html.NewTokenizer(body)
	defer body.Close()
	ulCount := 0

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return events
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == "ul" {
				ulCount++
			}

			if t.Data == "li" && ulCount == 2 {
				for {
					u := z.Next()
					ntext := z.Text()
					if u == html.TextToken {
						text = append(text, ntext...)
					}

					tk := z.Token()
					if tk.Data == "ul" && u == html.EndTagToken {
						result := string(text)
						temp := strings.Split(result, "\n")

						for _, eventText := range temp {
							event := fmt.Sprintf("On this day in %s", eventText)
							events = append(events, event)
						}
						return events
					}
				}
			}
		}
	}
	return events
}

func randomEvent(events []string) string {
	totalEvents := len(events) - 2
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(totalEvents)
	return events[r]
}

func main() {
	events := events()
	event := randomEvent(events)
	println(event)
}
