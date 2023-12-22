package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/rating/internal/models"
	"middleware/rating/internal/repositories/ratings"
	"net/http"
)

// UpdateUser
// @Tags         user
// @Summary      Update a user.
// @Description  Update the user with the specified ID.
// @Param  user  body 	  models.User{Name, Username} 	true  "User object to be updated"
// @Param  id    path     string  		true  "User UUID formatted ID"
// @Success 200  {object} models.User         "User object updated"
// @Failure 400  {object} models.CustomError  "Invalid JSON format"
// @Failure 404  {object} models.CustomError  "User not found"
// @Failure 422  {object} models.CustomError  "Cannot parse id"
// @Failure 500  {object} models.CustomError  "Something went wrong"
// @Router       /users/{id} 	[put]
func UpdateRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ratingId, _ := ctx.Value("ratingId").(uuid.UUID)
	var rating models.Rating
	err := json.NewDecoder(r.Body).Decode(&rating)

	ratingNew, err := ratings.UpdateRating(rating, ratingId)
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
	body, _ := json.Marshal(ratingNew)
	_, _ = w.Write(body)
	return
}
