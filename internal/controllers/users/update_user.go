package users

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/user/internal/models"
	"middleware/user/internal/repositories/users"
	"net/http"
)

// UpdateUser
// @Tags 		users
// @Summary 	Mettre à jour un utilisateur
// @Description Met à jour l'utilisateur avec l'ID spécifié.
// @Param 		id path string true "ID de l'utilisateur au format UUID"
// @Param 		user body models.User true "Objet utilisateur à mettre à jour"
// @Success 	200 {object} models.User "Utilisateur mis à jour"
// @Failure 	404 NOT FOUND
// @Failure      422            "Cannot parse id"
// @Failure 	500 {object} models.CustomError "Erreur interne du serveur"
// @Router 		/users/{id} [put]

// UpdateUser
// @Tags         user
// @Summary      Update a user.
// @Description  Update the user with the specified ID.
// @Param 		 user 			body 	  models.User 	true  "User object to be updated"
// @Param        id           	path      string  		true  "User UUID formatted ID"
// @Success      200            {object}  models.User
// @Failure 	 404 		 	"User not found"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} 	[put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, _ := ctx.Value("userId").(uuid.UUID)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	userNew, err := users.UpdateUser(user, userId)
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
