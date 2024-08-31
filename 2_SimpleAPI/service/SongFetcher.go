package service

import (
	"io"
	"log"
	"net/http"
)

func FetchSong(query string) (string, error) {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("SongFetcher: ")
	client := &http.Client{}

	req, err := http.NewRequest("get", "https://loripsum.net/api/1/short", nil)
	if err != nil {
		log.Default().Print("Issue building request")
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if res != nil && res.StatusCode == 200 {
		b, _ := io.ReadAll(res.Body)
		return string(b), nil
	}
	return "", nil
}
