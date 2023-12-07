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

	rows, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		return nil, err
	}

	helpers.CloseDB(db)

	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Id, &data.Name)
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
	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func CreateUser(user models.User) (*models.User, error) {

	randomUUID, err := uuid.NewV4()
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	
	_, err = db.Exec("INSERT INTO USERS (id, name) VALUES (?, ?)", randomUUID.String(), user.Name)
    if err != nil {
        return nil, err
    }

	row := db.QueryRow("SELECT * FROM USERS WHERE id = ?", randomUUID.String())
	helpers.CloseDB(db)

	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, err
	}
	
	return &user, err
}

func DeleteUser(id uuid.UUID) error {

	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM USERS WHERE id = ?", id.String())
	if err != nil {
		return err
	}

	helpers.CloseDB(db)

	return err
}

func UpdateUser(user models.User, id uuid.UUID) (*models.User, error) {
	
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("UPDATE USERS SET Name = ? WHERE ?", user.Name, id.String())
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM USERS WHERE id = ?", id.String())
	helpers.CloseDB(db)

	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, err
}
