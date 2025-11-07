package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"github.com/SiddhantSShende/bookings-app/pkg/config"
	"github.com/SiddhantSShende/bookings-app/pkg/handlers"
	"github.com/SiddhantSShende/bookings-app/pkg/render"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false // use true for production

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("Error creating template cache:", err)
		return
	}
	app.TemplateCache = tc
	app.UseCache = false // Set to true in production (false is for development mode)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting server on port %s", portNumber))

	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
