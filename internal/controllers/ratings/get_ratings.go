package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/rating/internal/models"
	service "middleware/rating/internal/services/ratings"
	"net/http"
)

// GetRatings
// @Tags         ratings
// @Summary      Get ratings
// @Description  Get all ratings.
// @Success 200  {array}  models.Rating 	 "Array of Rating object"
// @Failure 422  {object} models.CustomError "Cannot parse id"
// @Failure 500  {object} models.CustomError "Something went wrong"
// @Router       /songs/{song_id}/ratings/ 	[get]
func GetRatings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("songId").(uuid.UUID)

	ratings, err := service.GetAllRatings(songId)
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
	body, _ := json.Marshal(ratings)
	_, _ = w.Write(body)
	return
}
