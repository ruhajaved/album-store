package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM albums;")
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var albums album
		err := rows.Scan(&title)
		CheckError(err)
		fmt.Println(title)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	connStr := "postgresql://postgres@localhost/album_store?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	CheckError(err)

	// close database at the end
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)
	
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
