package Songs

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	service "middleware/example/internal/services/Songs"
	"net/http"
)

// CreateSong
// @Tags        songs
// @Summary     Create a new song
// @Description Create a new song with the provided parameters.
// @Param    song      body     models.SongRequest true "Song object to be created"
// @Success  201       {object} models.Song             "Created"
// @Failure  400       {object} models.CustomError      "Invalid request"
// @Failure  500       {object} models.CustomError      "Something went wrong"
// @Router   /songs/   [post]
func CreateSong(w http.ResponseWriter, r *http.Request) {
	var song models.SongRequest
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	err = models.ValidateSongRequest(song)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newSong, err := service.CreateSong(song)
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
	body, _ := json.Marshal(newSong)
	_, _ = w.Write(body)
	return
}
