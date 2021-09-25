package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {

	//myChannel := make(chan float64)
	//myChannel <- 3.14
	//q := <- myChannel

	//sizes := make(chan int)
	//go responseSize("https://golang.org/", sizes)
	//go responseSize("https://golang.org/doc", sizes)
	//
	//fmt.Println(<- sizes)
	//fmt.Println(<- sizes)

	pages := make(chan Page)
	urls := []string{
		"https://golang.org/",
		"https://golang.org/doc",
		"https://example.com/",
		"http://yandex.ru/",
		"http://youtube.com/",
	}

	for _, url := range urls {
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Println(page.URL, " ", page.Size)
	}
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}
