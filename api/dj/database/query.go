package database

import (
	"server/api/dj/database/model"

	"gorm.io/gorm"
)

func GetQuery(db *gorm.DB, name string) (model.Query, error) {
	query := model.Query{}
	res := db.Where("name = ?", name).First(&query)
	if res.Error != nil {
		return query, res.Error
	}

	return query, nil
}
