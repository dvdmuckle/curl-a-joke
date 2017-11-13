package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Joke struct {
	ID   uint `gorm:"primary_key"`
	Joke string
}

func randjoke(dbFile string) (joke string) {
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

func requestjoke(w http.ResponseWriter, r *http.Request, dbFile string) {
	fmt.Fprintln(w, randjoke(dbFile))
}

func setup(db *string, port *int) (dbFile string, jokePort int) {
	dbFile = *db
	jokePort = *port
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return dbFile, jokePort
}

func main() {
	dbPtr := flag.String("jokesdb", "jokes.db", "Location to the jokes database")
	portPtr := flag.Int("port", 8080, "Port for server")
	flag.Parse()
	dbFile, jokePort := setup(dbPtr, portPtr)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestjoke(w, r, dbFile)
	})
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(jokePort),
		handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
