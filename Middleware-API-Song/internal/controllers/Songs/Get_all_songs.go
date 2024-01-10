package Songs

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	service "middleware/example/internal/services/Songs"
	"net/http"
)

// GetAllSongs
// @Tags         songs
// @Summary      Get songs.
// @Description  Get all songs.
// @Success      200       {array}    models.Song              "Array of Song object"
// @Failure      500       {object}   models.CustomError       "Something went wrong"
// @Router       /songs/      [get]
func GetAllSongs(w http.ResponseWriter, _ *http.Request) {
	songs, err := service.GetAllSongs()
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
	body, _ := json.Marshal(songs)
	_, _ = w.Write(body)
	return
}
