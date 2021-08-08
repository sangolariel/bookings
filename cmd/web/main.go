package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sangolariel/bookings/internal/config"
	"github.com/sangolariel/bookings/internal/driver"
	"github.com/sangolariel/bookings/internal/handlers"
	"github.com/sangolariel/bookings/internal/helpers"
	"github.com/sangolariel/bookings/internal/models"
	"github.com/sangolariel/bookings/internal/render"
)

const port = ":8080"

var app config.AppConfig

var session *scs.SessionManager

var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	db, err := run()

	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	fmt.Printf("Aplication open on port %s", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	gob.Register(models.Reservation{})
	//Env
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//connect to the db
	log.Println("Connecting to Database ....")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=sangnguyen password=")
	if err != nil {
		log.Fatal("Can not conect to Database.")
	}

	log.Println("Connected to database")

	tc, err := render.CreateTemplateCatche()
	if err != nil {
		log.Fatal("Can't create Template catche")
		return nil, err
	}
	app.TemplateCatche = tc
	app.UseCatche = false

	repo := handlers.NewRepository(&app, db)
	handlers.NewHandler(repo)
	helpers.NewHelpers(&app)

	render.NewTemplates(&app)

	return db, nil
}
