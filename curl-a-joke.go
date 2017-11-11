package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"math/rand"
	"net/http"
	"os"
)

type Joke struct {
	ID   uint `gorm:"primary_key"`
	Joke string
}

var dbFile string
var jokePort string

func randjoke() (joke string) {
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "Error!"
	}
	defer db.Close()
	jokes := []Joke{}
	db.Find(&jokes)
	j := jokes[rand.Intn(len(jokes))]
	return j.Joke
}
func requestjoke(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, randjoke())
}
func main() {
	dbPtr := flag.String("jokesdb", "jokes.db", "Location to the jokes database")
	portPtr := flag.Int("port", 8080, "Port for server")
	flag.Parse()
	dbFile = *dbPtr
	jokePort = fmt.Sprintf(":%d", *portPtr)
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	http.HandleFunc("/", requestjoke)
	err := http.ListenAndServe(jokePort, nil)
	fmt.Fprintln(os.Stderr, err)
}
