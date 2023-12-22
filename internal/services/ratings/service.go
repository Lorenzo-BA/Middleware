package ratings

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/rating/internal/models"
	repository "middleware/rating/internal/repositories/ratings"
	"net/http"
)

func GetAllRatings(songId uuid.UUID) ([]models.Rating, error) {
	ratings, err := repository.GetAllRatings(songId)
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return ratings, err
}

func GetRatingById(songId uuid.UUID, ratingId uuid.UUID) (*models.Rating, error) {
	rating, err := repository.GetRatingById(songId, ratingId)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "rating not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return rating, err
}

func CreateRating(rating models.Rating, songId uuid.UUID) (*models.Rating, error) {
	createdRating, err := repository.CreateRating(rating, songId)
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return createdRating, err
}

func DeleteRating(songId uuid.UUID, ratingId uuid.UUID) error {
	err := repository.DeleteRating(songId, ratingId)
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return err
}

func UpdateRating(rating models.Rating, songId uuid.UUID, ratingId uuid.UUID) (*models.Rating, error) {
	UpdatedRating, err := repository.UpdateRating(rating, songId, ratingId)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "rating not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return UpdatedRating, err
}
