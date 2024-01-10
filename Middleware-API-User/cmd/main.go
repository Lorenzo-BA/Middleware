package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/user/internal/controllers/users"
	"middleware/user/internal/helpers"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetUsers)
		r.Post("/", users.CreateUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(users.Ctx)
			r.Get("/", users.GetUser)
			r.Delete("/", users.DeleteUser)
			r.Put("/", users.UpdateUser)
		})
	})

	port := ":8080"
	logrus.Info("[INFO] Web server started. Now listening on %s", port)
	logrus.Fatalln(http.ListenAndServe(port, r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS USERS (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			inscription_date TIMESTAMP NOT NULL,
			name VARCHAR(255) NOT NULL,
			username VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
