package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/rating/internal/models"
	"middleware/rating/internal/repositories/ratings"
	"net/http"
)

// CreateUser
// @Tags 		users
// @Summary 	Create a new user
// @Description Create a new user with the provided name.
// @Param  user  body 	 models.User  true  "User object to be created"
// @Success 201 {object} models.User        "Created"
// @Failure 400 {object} models.CustomError "Invalid JSON format"
// @Failure 500 {object} models.CustomError	"Something went wrong"
// @Router 		/users/ 			[post]
func CreateRating(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	ratingId, _ := ctx.Value("ratingId").(uuid.UUID)

	newRating, err := ratings.CreateRating(rating, ratingId)
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
	body, _ := json.Marshal(newRating)
	_, _ = w.Write(body)
	return
}
