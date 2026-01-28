package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	_, err := File_writer("urls.txt", "http://www.test.com")

	if err != nil {
		log.Fatal(err)
	}

	list, err := File_reader("urls.txt")

	if err != nil {
		log.Fatal(err)
	}

	for _, url := range list {
		res, err := http.Get(url)

		if err != nil {
			fmt.Printf("%s is down (Error: %b\n", url, err)
			continue

		}

		if res.StatusCode == http.StatusOK {
			log.Printf("%s is up!\n", url)
		}

		res.Body.Close()

	}

}
