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
	_, err = db.Exec("INSERT INTO SONGS (id, content) VALUES ('6ba7b810-9dad-11d1-80b4-00c04fd430c8','NewSong')")
	rows, err := db.Query("SELECT * FROM Songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Content)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
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
	err = row.Scan(&song.Id, &song.Content)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func PostSongById(song models.Song) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	randomUUID, err := uuid.NewV4()
	_, err = db.Exec("INSERT INTO Songs (id, content) values (?, ?)", randomUUID.String(), song.Content)
	if err != nil {
		return nil, err
	}

	PostSong := &models.Song{Id: &randomUUID, Content: song.Content}
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}
	return PostSong, err
}

func PutSongById(song models.Song) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	randomUUID, err := uuid.NewV4()
	_, _ = db.Exec("INSERT INTO Songs (id, content) values (?, ?)", randomUUID.String(), song.Content)
	PostSong := &models.Song{Id: &randomUUID, Content: song.Content}
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}
	return PostSong, err
}

func DeleteSongById(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	fmt.Print("tset5 \n")
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

	fmt.Print("tset6 \n")

	return err
}
