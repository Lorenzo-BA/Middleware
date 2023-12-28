package Songs

import (
	"fmt"
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM Songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Title, &data.FileName, &data.Artist, &data.PublishedDate)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}

	_ = rows.Close()
	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM Songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Title, &song.FileName, &song.Artist, &song.PublishedDate)
	if err != nil {
		return nil, err
	}

	return &song, err
}

func CreateSong(song models.SongRequest) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	randomUUID, _ := uuid.NewV4()

	_, err = db.Exec(
		"INSERT INTO Songs (id, title, file_name, artist,  published_date) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)",
		randomUUID.String(),
		song.Title,
		song.FileName,
		song.Artist,
	)
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	newSong, err := GetSongById(randomUUID)
	return newSong, err
}

func UpdateSong(song models.SongRequest, id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	fmt.Print(id)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("UPDATE Songs SET title = ?, file_name = ?, artist = ?  WHERE id = ?", song.Title, song.FileName, song.Artist, id.String())
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}
	fmt.Print(id)

	newSong, err := GetSongById(id)
	return newSong, err
}

func DeleteSong(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM Songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	return err
}
