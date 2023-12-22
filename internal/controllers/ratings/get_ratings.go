package ratings

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/rating/internal/models"
	"middleware/rating/internal/services/ratings"
	"net/http"
)

// GetUsers
// @Tags         users
// @Summary      Get users.
// @Description  Get all users.
// @Success 200 {array}  models.User 		"Array of User object"
// @Failure 500 {object} models.CustomError "Something went wrong"
// @Router      /users/ 		[get]
func GetRatings(w http.ResponseWriter, _ *http.Request) {
	ratings, err := ratings.GetAllRatings()
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
