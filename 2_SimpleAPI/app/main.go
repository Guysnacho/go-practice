// We'll be building to the beat of this drum https://gowebexamples.com/hello-world/
package main

import (
	"fmt"
	"log"
	"net/http"

	fetcher "tunjiprod.com/service"
	util "tunjiprod.com/util"
)

func fetchSong(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		log.Default().Println("No request recieved")
		http.Error(w, "No request recieved", 401)
		return
	}
	body, err := util.ParseBody(&r.Body)
	if err != nil {
		log.Default().Println("Issue parsing")
		http.Error(w, "Issue parsing", 401)
		return
	}
	fmt.Printf("body: %v\n", body)

	b, _ := fetcher.FetchSong("Tunji")
	log.Default().Println(b)
	fmt.Fprintf(w, "Something something, here's lorem :)\n%v", b)
}

func main() {
	http.HandleFunc("/", fetchSong)
	http.ListenAndServe(":8080", nil)
}
