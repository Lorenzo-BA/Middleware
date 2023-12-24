package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/rating/internal/models"
	service "middleware/rating/internal/services/ratings"
	"net/http"
)

// DeleteRating
// @Tags      ratings
// @Summary   Delete a rating
// @Description  Delete the rating with the specified ID.
// @Param     songId    path    string            true   "Song UUID formatted ID"
// @Param     ratingId  path    string            true   "Rating UUID formatted ID"
// @Success   204                                        "No Content"
// @Failure   422       {object} models.CustomError      "Cannot parse id"
// @Failure   500       {object} models.CustomError      "Something went wrong"
// @Router    /songs/{song_id}/ratings/{rating_id}       [delete]
func DeleteRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("songId").(uuid.UUID)
	ratingId, _ := ctx.Value("ratingId").(uuid.UUID)

	err := service.DeleteRating(songId, ratingId)
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

	w.WriteHeader(http.StatusNoContent)
	return
}
