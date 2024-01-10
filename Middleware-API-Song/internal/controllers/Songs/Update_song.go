package Songs

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	service "middleware/example/internal/services/Songs"
	"net/http"
)

// UpdateSong
// @Tags         songs
// @Summary      Update a song.
// @Description  Update the song with the specified ID.
// @Param        song      body       models.SongRequest true  "Song object to be updated"
// @Param        id        path       string             true  "Song UUID formatted ID"
// @Success      200       {object}   models.Song              "Song object updated"
// @Failure      400       {object}   models.CustomError       "Invalid request"
// @Failure      404       {object}   models.CustomError       "Song not found"
// @Failure      422       {object}   models.CustomError       "Cannot parse id"
// @Failure      500       {object}   models.CustomError       "Something went wrong"
// @Router       /songs/{id}      [put]
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("songId").(uuid.UUID)
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

	newSong, err := service.UpdateSong(song, songId)
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
	body, _ := json.Marshal(newSong)
	_, _ = w.Write(body)
	return
}
