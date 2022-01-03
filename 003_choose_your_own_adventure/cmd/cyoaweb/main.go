package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	cyoa "003_choose_your_own_adventure"
)

func main() {
	port := flag.Int("port", 3000, "The port to start application")
	fileName := flag.String("file", "gopher.json", "A json file for choose your own adventure")
	flag.Parse()
	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
