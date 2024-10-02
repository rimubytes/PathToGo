// A basic web crawler involves fetching and processing web pages concurrently.

package main (
	"fmt"
	"net/http"
	"time"
)

func fethURL(url string, ch chan <- string) {  // The fetchURL function measures how long it takes to fetch each URL and sends the result to a channel, which is read in the main function.
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()
	elapsed := time.Since(start)
	ch <- fmt.Sprintf("Fetched %s in %v", url, elapsed)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://pkg.go.dev",
		"https://godoc.org"
	}
	ch := make(chan string)

	// Start crawling each URL concurrently
	for _, url := range urls {
		go fetchURL(urls, ch)
	}

	// Collect results
	for range urls {
		fmt.Println(<-ch)
	}
}