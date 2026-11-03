package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetch(url string) (int, time.Duration) {
	start := time.Now()
	resp, err := http.Get(url)
	elapsed := time.Since(start)
	if err != nil {
		return 0, elapsed
	}
	defer resp.Body.Close()
	return resp.StatusCode, elapsed


}

func main() {

	urls := []string{
		// Reliable, fast sites
		"https://www.google.com",
		"https://www.github.com",
		"https://go.dev",
		"https://pkg.go.dev",
		"https://www.wikipedia.org",
		"https://www.mozilla.org",
		"https://www.cloudflare.com",
		"https://www.apache.org",
		"https://www.python.org",
		"https://www.rust-lang.org",
		"https://news.ycombinator.com",
		"https://www.kernel.org",
		"https://www.debian.org",
		"https://www.gnu.org",
		"https://stackoverflow.com",

		// httpbin: useful for testing behaviors
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404", // not found
		"https://httpbin.org/status/500", // server error
		"https://httpbin.org/redirect/2", // follows 2 redirects
		"https://httpbin.org/delay/3",    // slow response (3s)
	}

	start := time.Now()

	
	// https://go.dev                     status=200  took=210ms
	for _, url := range urls {
		status, elapsed := fetch(url)
		fmt.Println(url, "status = ", status, "took", elapsed)

	}

	fmt.Println("total:", time.Since(start))
}
