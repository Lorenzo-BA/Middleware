package users

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/user/internal/models"
	service "middleware/user/internal/services/users"
	"net/http"
)

// DeleteUser
// @Tags        users
// @Summary     Delete a user
// @Description Delete the user with the specified ID.
// @Param       id       path       string      true   "User UUID formatted ID"
// @Success     204      {object}   models.User        "No Content"
// @Failure     422      {object}   models.CustomError "Cannot parse id"
// @Failure     500      {object}   models.CustomError "Something went wrong"
// @Router      /users/{id}      [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, _ := ctx.Value("userId").(uuid.UUID)

	err := service.DeleteUser(userId)
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

	w.WriteHeader(http.StatusNoContent)
	return
}