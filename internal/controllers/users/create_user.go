package users

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/users"
	"net/http"
)

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

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(newUser)
	_, _ = w.Write(body)
	return
}
