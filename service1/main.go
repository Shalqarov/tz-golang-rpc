package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var letterRunes = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

type Salt struct {
	Salt string `json:"salt"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Generate() string {
	salt := make([]rune, 12)
	for i := range salt {
		salt[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(salt)
}

func main() {
	port := flag.String("addr", ":8080", `Network Port
	USAGE:
	--addr=:8080`)
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/generate-salt", func(w http.ResponseWriter, r *http.Request) {
		s := &Salt{Generate()}
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(s); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	log.Printf("generate-salt started at %s\n", *port)
	if err := http.ListenAndServe(*port, r); err != nil {
		log.Fatalln(err)
	}
}
