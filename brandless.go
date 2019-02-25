package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/chriskaukis/brandless/icanhazdadjoke"
	"github.com/chriskaukis/brandless/markov"
)

var (
	markovPrefixLength = 1
	defaultPort        = "8080"
	templates          = template.Must(template.ParseFiles("templates/index.html"))
)

func main() {
	// Required to get a decent Markov message.
	rand.Seed(time.Now().UTC().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", IndexHandler)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}

type IndexPage struct {
	Joke       string
	MarkovJoke string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	joke, err := getRandomJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	page := &IndexPage{
		Joke:       joke.Joke,
		MarkovJoke: markovIt(joke.Joke),
	}
	err = templates.ExecuteTemplate(w, "index.html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getRandomJoke() (*icanhazdadjoke.Joke, error) {
	client := icanhazdadjoke.New()
	return client.Random()
}

func markovIt(s string) string {
	markov := markov.New(markovPrefixLength)
	markov.Build(s)
	return markov.Generate(len(s))
}
