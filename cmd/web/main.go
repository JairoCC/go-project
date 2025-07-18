package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JairoCC/go-project/pkg/config"
	"github.com/JairoCC/go-project/pkg/handlers"
	"github.com/JairoCC/go-project/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	
	fmt.Println(fmt.Sprintf("starint application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
