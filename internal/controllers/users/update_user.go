package users

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/users"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
	ctx := r.Context()
	userId, _ := ctx.Value("userId").(uuid.UUID)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	var userNew *models.User
	userNew, err = users.UpdateUserById(user, userId)
	
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
	body, _ := json.Marshal(userNew)
	_, _ = w.Write(body)
	return
}

