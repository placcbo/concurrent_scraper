package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://go.dev",
		"https://github.com",
		"https://golang.org",
	}


	for _, url := range urls{
resp, err := http.Get(url)
fmt.Println(resp.StatusCode, err)
	}

	

	

}
