package users

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/user/internal/models"
	service "middleware/user/internal/services/users"
	"net/http"
)

// CreateUser
// @Tags        users
// @Summary     Create a new user
// @Description Create a new user with the provided name.
// @Param    user      body     models.UserRequest true "User object to be created"
// @Success  201       {object} models.User             "Created"
// @Failure  400       {object} models.CustomError      "Invalid request"
// @Failure  500       {object} models.CustomError      "Something went wrong"
// @Router   /users/   [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	err = models.ValidateUserRequest(user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newUser, err := service.CreateUser(user)
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
