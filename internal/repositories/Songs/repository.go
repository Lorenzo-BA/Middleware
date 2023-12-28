package Songs

import (
	"fmt"
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"time"
)

func RequestGetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM Songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Content, &data.Title, &data.File_name, &data.Artist, &data.Published_date)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}

func RequestGetSong(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM Songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Content, &song.Title, &song.File_name, &song.Artist, &song.Published_date)

	if err != nil {
		return nil, err
	}
	return &song, err
}

func RequestCreateSong(song models.Song) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	randomUUID, err := uuid.NewV4()
	_, err = db.Exec(
		"INSERT INTO Songs (id, content, title, file_name, artist,  Published_date) VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP)",
		randomUUID.String(),
		song.Content,
		song.Title,
		song.File_name,
		song.Artist,
	)
	if err != nil {
		return nil, err
	}

	PostSong := &models.Song{Id: &randomUUID, Content: song.Content, Title: song.Title, Artist: song.Artist, File_name: song.File_name, Published_date: time.Now()}
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}
	return PostSong, err
}

func RequestUpgradeSong(song models.Song, id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("UPDATE Songs SET content = ?, title = ?, file_name = ?, artist = ?,  WHERE id = ?", song.Content, song.Title, song.File_name, song.Artist, id.String())
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM Songs WHERE id = ?", id.String())
	helpers.CloseDB(db)
	err = row.Scan(&song.Id, &song.Content, &song.Title, &song.File_name, &song.Artist, &song.Published_date)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func RequestDeleteSong(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	//_, err = uuid.NewV4()
	_, err = db.Exec("DELETE FROM Songs WHERE id=?", id.String())
	if err != nil {
		fmt.Println("Error deleting song:", err)
		return err
	}

	helpers.CloseDB(db)

	if err != nil {
		return err
	}

	return err
}
