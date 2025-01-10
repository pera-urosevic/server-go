package database

import "server/api/gallery/types"

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
