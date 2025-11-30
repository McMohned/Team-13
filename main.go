package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func downloadURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR downloading %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR reading response from %s: %v\n", url, err)
		return
	}

	fmt.Printf("URL: %s | Status: %d | Size: %d bytes\n", url, resp.StatusCode, len(body))
}

func main() {
	// List of URLs to download
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.wikipedia.org",
		"https://www.reddit.com",
	}

	var wg sync.WaitGroup

	fmt.Println("Starting downloads...\n")

	for _, url := range urls {
		wg.Add(1)
		go downloadURL(url, &wg) // Launch goroutine
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("\nAll downloads are complete.")
}
