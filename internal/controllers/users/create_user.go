package users

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/user/internal/models"
	"middleware/user/internal/repositories/users"
	"net/http"
)

// CreateUser
// @Tags 		users
// @Summary 	Create a new user
// @Description Create a new user with the provided name.
// @Param 		user 			body 		models.User 	true 	"User object to be created"
// @Success 	201 			{object} 	models.User
// @Failure 	500 			"Something went wrong"
// @Router 		/users 			[post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser, err := users.CreateUser(user)
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
	body, _ := json.Marshal(newUser)
	_, _ = w.Write(body)
	return
}
