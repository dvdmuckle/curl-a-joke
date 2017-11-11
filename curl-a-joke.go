package main

import (
	"flag"
	"fmt"
	_ "github.com/jinzhu/gorm"
	"net/http"
)

type Joke struct {
	ID   uint `gorm:"primary_key"`
	Joke string
}

var dbFile string
var jokePort string

func requestjoke(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "There should be a joke here. Not yet, but soon.")
}
func main() {
	dbPtr := flag.String("jokedb", "jokes.db", "Location to the jokes database")
	portPtr := flag.Int("port", 8080, "Port for server")
	flag.Parse()
	dbFile = *dbPtr
	jokePort = fmt.Sprintf(":%d", *portPtr)
	http.HandleFunc("/", requestjoke)
	http.ListenAndServe(jokePort, nil)
}
