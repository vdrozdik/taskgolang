package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	value := <-awaitTask()

	fmt.Println(value)
}
func awaitTask() <-chan string {
	fmt.Println("Starting Task...")

	c := make(chan string)

	go func() {
		resp, err := http.Get("https://go.dev")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		c <- string(body)

		fmt.Println("...Done!")
	}()

	return c
}
