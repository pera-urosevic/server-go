package database

import (
	"os"
	"server/api/gallery/types"

	"database/sql"

	_ "modernc.org/sqlite"
)

func Database() (*sql.DB, error) {
	dbPath := os.Getenv("GALLERY_DB_PATH")
	db, err := sql.Open("sqlite", dbPath)
	return db, err
}

func GetPhotos(filter string) ([]types.Photo, error) {
	db, err := Database()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	f := "%" + filter + "%"
	rows, err := db.Query("SELECT * FROM [gallery] WHERE [album] LIKE ? OR [datetime] LIKE ? OR [title] LIKE ? OR [description] LIKE ? OR [keywords] LIKE ? ORDER BY [datetime] DESC", f, f, f, f, f)
	if err != nil {
		return nil, err
	}
	photos := []types.Photo{}
	for rows.Next() {
		photo := types.Photo{}
		err := rows.Scan(&photo.ID, &photo.Path, &photo.Type, &photo.Modified, &photo.Online, &photo.Album, &photo.Datetime, &photo.Title, &photo.Description, &photo.Keywords, &photo.Copyright)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func GetPhoto(photoID int64) (types.Photo, error) {
	db, err := Database()
	if err != nil {
		return types.Photo{}, err
	}
	defer db.Close()
	photo := types.Photo{}
	row := db.QueryRow("SELECT * FROM [gallery] WHERE [id] = ?", photoID)
	err = row.Scan(&photo.ID, &photo.Path, &photo.Type, &photo.Modified, &photo.Online, &photo.Album, &photo.Datetime, &photo.Title, &photo.Description, &photo.Keywords, &photo.Copyright)
	if err != nil {
		return types.Photo{}, err
	}
	return photo, nil
}

func InsertPhoto(photo types.Photo) (int64, error) {
	db, err := Database()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	res, err := db.Exec("INSERT INTO [gallery] ([path], [type], [modified], [online], [album], [datetime], [title], [description], [keywords], [copyright]) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", photo.Path, photo.Type, photo.Modified, photo.Online, photo.Album, photo.Datetime, photo.Title, photo.Description, photo.Keywords, photo.Copyright)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdatePhoto(photo types.Photo) error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("UPDATE [gallery] SET [path]=?, [type]=?, [modified]=?, [online]=?, [album]=?, [datetime]=?, [title]=?, [description]=?, [keywords]=?, [copyright]=? WHERE [id] = ?", photo.Path, photo.Type, photo.Modified, photo.Online, photo.Album, photo.Datetime, photo.Title, photo.Description, photo.Keywords, photo.Copyright, photo.ID)
	if err != nil {
		return err
	}
	return nil
}

func RemovePhoto(postID int64) error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM [gallery] WHERE [id] = ?", postID)
	if err != nil {
		return err
	}
	return nil
}
