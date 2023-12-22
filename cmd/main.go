package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/rating/internal/controllers/ratings"
	"middleware/rating/internal/helpers"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Route("/songs/{song_id}/ratings", func(r chi.Router) {
		r.Use(ratings.CtxSongId)
		r.Get("/", ratings.GetRatings)
		r.Post("/", ratings.CreateRating)
		r.Route("/{rating_id}", func(r chi.Router) {
			r.Use(ratings.CtxRatingId)
			r.Get("/", ratings.GetRating)
			r.Delete("/", ratings.DeleteRating)
			r.Put("/", ratings.UpdateRating)
		})
	})

	port := ":8081"
	logrus.Info("[INFO] Web server started. Now listening on %s", port)
	logrus.Fatalln(http.ListenAndServe(port, r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS RATINGS (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			rating_date TIMESTAMP NOT NULL,
			comment VARCHAR(255) NOT NULL,
			song_id VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			rating INT NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
