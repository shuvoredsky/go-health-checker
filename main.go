package main

import (
    "fmt"
    "net/http"
    "time"
)


func main(){
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
	}

	start := time.Now()

	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(url, "is down")
		} else {
			fmt.Println(url, "is up")
			res.Body.Close()
		}
	}

	fmt.Println("time taken:", time.Since(start))
	fmt.Println("All url checked successfully")
}