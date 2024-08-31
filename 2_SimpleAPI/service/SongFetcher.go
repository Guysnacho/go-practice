package service

import (
	"log"
	"net/http"
)

func FetchSong(query string) (string, error) {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("SongFetcher: ")

	res, err := http.Get("https://loripsum.net/api/1/short")
	if err != nil {
		log.Default().Print("Issue requesting data")
		return "", err
	}
	var data []byte

	log.Default().Println("Response recieved")
	log.Default().Println(res)

	defer res.Body.Close()

	if res != nil && res.StatusCode == 200 {
		_, err := res.Body.Read(data)
		if err != nil {
			return "", err
		} else {
			log.Default().Println("Data recieved")
			log.Default().Println(data)

			return string(data), nil
		}
	}
	return "", nil
}
