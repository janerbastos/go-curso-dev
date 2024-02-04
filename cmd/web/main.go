package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/janerbastos/go-curso/pkg/config"
	"github.com/janerbastos/go-curso/pkg/handlers"
	"github.com/janerbastos/go-curso/pkg/hender"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := hender.CraateTemplateCache()
	if err != nil {
		log.Fatal("connot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	hender.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/abaut/", handlers.Repo.Abaut)
	fmt.Println(fmt.Sprintf("Start application in port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
