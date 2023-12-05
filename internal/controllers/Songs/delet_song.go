package Songs

import (
	"encoding/json"
	"fmt"
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
	fmt.Print("tset0000011110 \n")
	err := Songs.DeleteSongById(Song_Id)
	if err != nil {
		fmt.Print("tset")
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			fmt.Print("tset3")
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			fmt.Print("tset4")
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Print("tset2")

		return
	}

	fmt.Print("tset000000 \n")
	w.WriteHeader(http.StatusOK)
	return
}
