package ratings

import (
	"github.com/gofrs/uuid"
	"middleware/rating/internal/helpers"
	"middleware/rating/internal/models"
	"time"
)

func GetAllRatings(songId uuid.UUID) ([]models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM RATINGS WHERE song_id = ?", songId.String())
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	ratings := []models.Rating{}
	for rows.Next() {
		var data models.Rating
		err = rows.Scan(&data.Id, &data.RatingDate, &data.Comment, &data.SongId, &data.UserId, &data.Rating)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, data)
	}

	_ = rows.Close()
	return ratings, err
}

func GetRatingById(songId uuid.UUID, ratingId uuid.UUID) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM RATINGS WHERE id=? AND song_id=?", ratingId.String(), songId.String())
	helpers.CloseDB(db)

	var rating models.Rating
	err = row.Scan(&rating.Id, &rating.RatingDate, &rating.Comment, &rating.SongId, &rating.UserId, &rating.Rating)
	if err != nil {
		return nil, err
	}

	return &rating, err
}

func CreateRating(rating models.RatingCreateRequest, songId uuid.UUID) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	randomUUID, err := uuid.NewV4()
	if err != nil {
		helpers.CloseDB(db)
		return nil, err
	}

	ratingDate := time.Now()
	_, err = db.Exec("INSERT INTO RATINGS (id, rating_date, comment, song_id, user_id, rating) VALUES (?, ?, ?, ?, ?, ?)", randomUUID.String(), ratingDate, rating.Comment, songId, rating.UserId, rating.Rating)
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	createdRating, err := GetRatingById(songId, randomUUID)
	return createdRating, err
}

func DeleteRating(songId uuid.UUID, ratingId uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM RATINGS WHERE id = ? AND song_id = ?", ratingId.String(), songId.String())
	helpers.CloseDB(db)
	if err != nil {
		return err
	}

	return err
}

func UpdateRating(rating models.RatingUpdateRequest, songId uuid.UUID, ratingId uuid.UUID) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("UPDATE RATINGS SET comment = ?, rating = ? WHERE id = ? AND song_id = ?", rating.Comment, rating.Rating, ratingId.String(), songId.String())
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	updatedRating, err := GetRatingById(songId, ratingId)
	return updatedRating, err
}
