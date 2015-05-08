package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func worker(ticker <-chan time.Time) {
	for range ticker {
		resp, err := http.Get("http://web_1:5000/")
		if err != nil {
			log.Print("error connecting: ", err)
		}
		defer resp.Body.Close()

		var seen int
		fmt.Fscanf(
			resp.Body,
			"Hello World! I have been seen %d times.\n",
			&seen,
		)
		fmt.Printf("Seen %d\n", seen)
	}
}

func main() {
	go worker(time.Tick(time.Second * 2))
	go worker(time.Tick(time.Second * 10))
	worker(time.Tick(time.Millisecond * 500))
}
