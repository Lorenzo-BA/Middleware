package users

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/user/internal/models"
	service "middleware/user/internal/services/users"
	"net/http"
)

// GetUser
// @Tags         users
// @Summary      Get a user.
// @Description  Get the user with the specified ID.
// @Param        id        path       string       true        "User UUID formatted ID"
// @Success      200       {object}   models.User              "User object"
// @Failure      404       {object}   models.CustomError       "User not found"
// @Failure      422       {object}   models.CustomError       "Cannot parse id"
// @Failure      500       {object}   models.CustomError       "Something went wrong"
// @Router       /users/{id}      [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, _ := ctx.Value("userId").(uuid.UUID)

	user, err := service.GetUserById(userId)
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
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
	return
}
