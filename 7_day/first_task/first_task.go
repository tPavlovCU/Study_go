package main

import (
	"fmt"
	"net/http"
	"sync"
)

func ping(wg *sync.WaitGroup, url string) {
	response, err := http.Get(url)
	defer wg.Done()

	if err != nil {
		fmt.Println(url, "Ошибка")
		return
	}

	defer response.Body.Close()

	fmt.Println("Успешно", url, response.StatusCode)
}

func main() {
	wg := sync.WaitGroup{}

	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://yandex.ru",
		"https://habbbbbbbbbbbbbbbbbbbbbbbbr.com",
		"https://golang.org",
	}

	for _, value := range urls {
		wg.Add(1)
		go ping(&wg, value)
	}
	wg.Wait()
}
