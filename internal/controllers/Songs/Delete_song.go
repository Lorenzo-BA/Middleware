package Songs

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	service "middleware/example/internal/services/Songs"
	"net/http"
)

// DeleteSong
// @Tags        songs
// @Summary     Delete a song
// @Description Delete the song with the specified ID.
// @Param       id       path       string      true   "Song UUID formatted ID"
// @Success     204      null                          "No Content"
// @Failure     422      {object}   models.CustomError "Cannot parse id"
// @Failure     500      {object}   models.CustomError "Something went wrong"
// @Router      /songs/{id}      [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("songId").(uuid.UUID)

	err := service.DeleteSong(songId)
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
	return
}
