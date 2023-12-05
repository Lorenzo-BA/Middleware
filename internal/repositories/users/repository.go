package users

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Id, &data.Content)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	// don't forget to close rows
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
	err = row.Scan(&user.Id, &user.Content)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func UpdateUserById(user models.User, id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("UPDATE USERS SET content = ? WHERE ?", user.Content, id)
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM USERS WHERE id = ?", id)
	helpers.CloseDB(db)

	err = row.Scan(&user.Id, &user.Content)
	if err != nil {
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

	_, err = db.Exec("INSERT INTO USERS (id, content) VALUES (?, ?)", randomUUID.String(), user.Content)
    if err != nil {
        return nil, err
    }

    createdUser := &models.User{Id: &randomUUID, Content: user.Content}
	helpers.CloseDB(db)

    return createdUser, nil
}

func DeleteUser(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("DELETE FROM USERS WHERE id = ?", id.String())
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.Id, &user.Content)
	if err != nil {
		return nil, err
	}
	return &user, err
}