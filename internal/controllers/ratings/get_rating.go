package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/rating/internal/models"
	service "middleware/rating/internal/services/ratings"
	"net/http"
)

// GetRating
// @Tags         ratings
// @Summary      Get a rating
// @Description  Get the rating with the specified ID.
// @Param   songId     path    string  true   "Song UUID formatted ID"
// @Param   ratingId   path    string  true   "Rating UUID formatted ID"
// @Success 200 {object} models.Rating		  "Rating object"
// @Failure 404 {object} models.CustomError   "Rating not found"
// @Failure 422 {object} models.CustomError   "Cannot parse id"
// @Failure 500 {object} models.CustomError   "Something went wrong"
// @Router      /songs/{song_id}/ratings/{rating_id} [get]
func GetRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("songId").(uuid.UUID)
	ratingId, _ := ctx.Value("ratingId").(uuid.UUID)

	rating, err := service.GetRatingById(songId, ratingId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(rating)
	_, _ = w.Write(body)
	return
}
