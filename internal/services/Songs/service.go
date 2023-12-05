package Songs

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"log"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/Songs"
	"net/http"
)

func GetAllSongs() ([]models.Song, error) {
	var err error
	// calling repository
	songs, err := repository.GetAllSongs()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving Songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetSongById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving Songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

func DeleteSongById(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	fmt.Print("tset5")
	//_, err = uuid.NewV4()
	_, err = db.Exec("DELETE FROM Songs WHERE id=?", id.String())
	if err != nil {
		log.Println("Error deleting song:", err)
		return err
	}

	helpers.CloseDB(db)

	if err != nil {
		return err
	}

	return err
}
