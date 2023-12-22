package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/rating/internal/models"
	service "middleware/rating/internal/services/ratings"
	"net/http"
)

// CreateRating
// @Tags 		ratings
// @Summary 	Create a new rating
// @Description Create a new rating with the provided content.
// @Param   rating body  models.Rating true "Rating object to be created"
// @Param   songId path  string        true "Song UUID formatted ID"
// @Success 201 {object} models.Rating      "Created"
// @Failure 400 {object} models.CustomError "Invalid JSON format"
// @Failure 422 {object} models.CustomError "Cannot parse id"
// @Failure 500 {object} models.CustomError	"Something went wrong"
// @Router 		/songs/{song_id}/ratings/ 	[post]
func CreateRating(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	songId, _ := ctx.Value("songId").(uuid.UUID)

	newRating, err := service.CreateRating(rating, songId)
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

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(newRating)
	_, _ = w.Write(body)
	return
}
