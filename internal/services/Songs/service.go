package Songs

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/Songs"
	"net/http"
)

func GetAllSongs() ([]models.Song, error) {
	songs, err := repository.GetAllSongs()
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
				Message: "Song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

func CreateSong(song models.SongRequest) (*models.Song, error) {
	newSong, err := repository.CreateSong(song)
	if err != nil {
		logrus.Errorf("error retrieving Songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return newSong, err
}

func UpdateSong(song models.SongRequest, id uuid.UUID) (*models.Song, error) {
	newSong, err := repository.UpdateSong(song, id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "Song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving Songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return newSong, err
}

func DeleteSong(id uuid.UUID) error {
	err := repository.DeleteSong(id)
	if err != nil {
		logrus.Errorf("error retrieving Songs : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return err
}
