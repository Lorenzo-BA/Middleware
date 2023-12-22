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

func GetAllRatings() ([]models.Rating, error) {
	ratings, err := repository.GetAllRatings()
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return ratings, err
}

func GetRatingById(id uuid.UUID) (*models.Rating, error) {
	rating, err := repository.GetRatingById(id)
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

func CreateRating(rating models.Rating, ratingId uuid.UUID) (*models.Rating, error) {
	newRating, err := repository.CreateRating(rating, ratingId)
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return newRating, err
}

func DeleteRating(id uuid.UUID) error {
	err := repository.DeleteRating(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return &models.CustomError{
				Message: "rating not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return err
}

func UpdateRating(rating models.Rating, id uuid.UUID) (*models.Rating, error) {
	newRating, err := repository.UpdateRating(rating, id)
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

	return newRating, err
}
