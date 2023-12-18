package users

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"middleware/user/internal/helpers"
	"middleware/user/internal/models"
	"net/http"
	"time"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM USERS")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Id, &data.InscriptionDate, &data.Name, &data.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}

	_ = rows.Close()
	return users, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM USERS WHERE id=?", id.String())
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.Id, &user.InscriptionDate, &user.Name, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.CustomError{
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
		return nil, err
	}

	return &user, err
}

func CreateUser(user models.User) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	randomUUID, err := uuid.NewV4()
	if err != nil {
		helpers.CloseDB(db)
		return nil, err
	}

	inscriptionDate := time.Now()
	_, err = db.Exec("INSERT INTO USERS (id, inscription_date, name, username) VALUES (?, ?, ?, ?)", randomUUID.String(), inscriptionDate, user.Name, user.Username)
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	createdUser, err := GetUserById(randomUUID)
	return createdUser, err
}

func DeleteUser(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM USERS WHERE id = ?", id.String())
	helpers.CloseDB(db)
	if err != nil {
		return err
	}

	return err
}

func UpdateUser(user models.User, id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("UPDATE USERS SET name = ?, username = ? WHERE id = ?", user.Username, user.Name, id.String())
	helpers.CloseDB(db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.CustomError{
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
		return nil, err
	}

	updatedUser, err := GetUserById(id)
	return updatedUser, err
}
