package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
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

func main() {
	//	 .xpath('//ul')[1]
	//	 .elements
	//	 .collect {|node| node.text.strip}
	body := eventsHtml()
	z := html.NewTokenizer(body)
	defer body.Close()
	ulCount := 0
	liCount := 0

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == "ul" {
				ulCount++
			}

			if t.Data == "li" && ulCount == 2 {

				for {
					dd := z.Next()
					dd = z.Next()

					switch {
					case dd == html.ErrorToken:
						return
					case dd == html.TextToken:
						println(string(z.Text()[:]))
					}
				}
				liCount++

				// z.Next()
				// ntext := z.Text()
				// text = append(text[:], ntext[:]...)

				// fmt.Println(string(text[:]))
			}
		}
	}
}
