package Songs

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/Songs"
	"net/http"
)

// GetAllSongs
// @Tags         Songs
// @Summary      Get Songs.
// @Description  Get Songs.
// @Success      200            {array}  models.Song
// @Failure      500             "Something went wrong"
// @Router       /Songs [get]
func GetAllSongs(w http.ResponseWriter, _ *http.Request) {
	// calling service
	Song, err := Songs.GetAllSongs()
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(Song)
	_, _ = w.Write(body)
	return
}
