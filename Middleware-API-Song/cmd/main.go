package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/Songs"
	"middleware/example/internal/helpers"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Route("/songs", func(r chi.Router) {
		r.Get("/", Songs.GetAllSongs)
		r.Post("/", Songs.CreateSong)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(Songs.Ctx)
			r.Get("/", Songs.GetSong)
			r.Delete("/", Songs.DeleteSong)
			r.Put("/", Songs.UpdateSong)
		})
	})

	port := ":8079"
	logrus.Info("[INFO] Web server started. Now listening on %s", port)
	logrus.Fatalln(http.ListenAndServe(port, r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS SONGS (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
    		title VARCHAR(255) NOT NULL,
        	file_name VARCHAR(255) NOT NULL,
    		artist VARCHAR(255) NOT NULL,
			published_date TIMESTAMP NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
