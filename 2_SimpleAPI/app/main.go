// We'll be building to the beat of this drum https://gowebexamples.com/hello-world/
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	fetcher "tunjiprod.com/service"
)

func fetchSong(c *gin.Context) {
	if c.Request.Method != "GET" {
		log.Default().Println("Invalid request")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// body, err := util.ParseBody(&r.Body)
	// if err != nil {
	// 	log.Default().Println("Issue parsing")
	// 	http.Error(w, "Issue parsing", 401)
	// 	return
	// }
	// fmt.Printf("body: %v\n", body)

	lorem, _ := fetcher.FetchSong("Tunji")
	log.Default().Println(lorem)
	fmt.Fprintf(c.Writer, "Something something, here's lorem :)%v", lorem)
}

func main() {
	router := gin.Default()
	router.GET("/", fetchSong)

	router.Run("localhost:8080")
}
