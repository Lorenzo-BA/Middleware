package Songs

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/Songs"
	"net/http"
)

// GetSong
// @Tags         Songs
// @Summary      Get a collection.
// @Description  Get a collection.
// @Param        id           	path      string  true  "Song UUID formatted ID"
// @Success      200            {object}  models.Song
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /Songs/{id} [get]
func DeletSong(w http.ResponseWriter, r *http.Request) {
	var _ models.Song
	ctx := r.Context()
	Song_Id, _ := ctx.Value("Song_Id").(uuid.UUID)
	err := Songs.DeleteSongById(Song_Id)
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
