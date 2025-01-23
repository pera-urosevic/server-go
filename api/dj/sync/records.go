package sync

import (
	"encoding/json"
	"server/api/dj/database/model"
	"server/api/dj/filesystem"
	"server/api/dj/log"

	"gorm.io/gorm"
)

func getRecords(db *gorm.DB) ([]model.Song, error) {
	records := []model.Song{}
	res := db.Find(&records)
	if res.Error != nil {
		log.Log(res.Error)
		return nil, res.Error
	}

	return records, nil
}

func removeRecords(db *gorm.DB, records []model.Song, files []model.Song) ([]model.Song, []string, error) {
	removed := []string{}
	foundRecords := []model.Song{}
	for _, record := range records {
		found := false
		for _, file := range files {
			if record.Path == file.Path && record.Datetime.Equal(file.Datetime) {
				found = true
				break
			}
		}
		if found {
			foundRecords = append(foundRecords, record)
		} else {
			removed = append(removed, record.Path)
			res := db.Where("path = ?", record.Path).Delete(&model.Song{})
			if res.Error != nil {
				log.Log(res.Error)
				return nil, nil, res.Error
			}
		}
	}
	return foundRecords, removed, nil
}

func addRecords(db *gorm.DB, records []model.Song, files []model.Song) ([]model.Song, []string, error) {
	added := []string{}
	for _, file := range files {

		found := false
		for _, record := range records {
			if record.Path == file.Path {
				found = true
				break
			}
		}

		if !found {
			added = append(added, file.Path)

			meta, err := filesystem.ReadMeta(file.Path)
			if err != nil {
				log.Log(err)
				return nil, nil, err
			}

			jsonMeta, err := json.Marshal(meta)
			if err != nil {
				log.Log(err)
				return nil, nil, err
			}

			song := model.Song{Path: file.Path, Meta: string(jsonMeta), Datetime: file.Datetime}
			res := db.Create(&song)
			if res.Error != nil {
				log.Log(res.Error)
				return nil, nil, res.Error
			}

			records = append(records, file)
		}
	}
	return records, added, nil
}
