package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/Songs"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/Music", func(r chi.Router) {
		r.Get("/", Songs.GetSongs)
		r.Post("/", Songs.PostSong)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(Songs.Ctx)
			r.Get("/", Songs.GetSong)
			r.Delete("/", Songs.DeletSong)
			r.Put("/", Songs.PutSong)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))

}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS Songs (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			content VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
