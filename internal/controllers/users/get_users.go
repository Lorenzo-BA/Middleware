package users

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/user/internal/models"
	"middleware/user/internal/services/users"
	"net/http"
)

// GetUsers
// @Tags         users
// @Summary      Get users.
// @Description  Get all users.
// @Success 200 {array}  models.User 		"Array of User object"
// @Failure 500 {object} models.CustomError "Something went wrong"
// @Router      /users/ 		[get]
func GetUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := users.GetAllUsers()
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
	body, _ := json.Marshal(users)
	_, _ = w.Write(body)
	return
}
