package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 	https://go.dev status=200 took=210ms
	// https://github.com status=200 took=340ms
	// https://golang.org status=200 took=180ms

	urls := []string{
		"https://go.dev",
		"https://github.com",
		"https://golang.org",
		"https://getdraftly.xyz",
	}

	for _, url := range urls {
		now := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("request failed", err)
			return
		}
		elapsed := time.Since(now)
		fmt.Println(url, resp.StatusCode, "took=", elapsed)
	}
}
