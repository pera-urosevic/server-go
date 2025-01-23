package controller

import (
	"server/api/cleaner/controller/lib"
	"server/api/cleaner/database"
	"server/api/cleaner/database/model"
	"server/system"

	"github.com/gin-gonic/gin"
)

func RemovePath(r *gin.Engine) {
	r.DELETE("/cleaner", func(c *gin.Context) {
		var record = model.Cleaner{}
		err := c.BindJSON(&record)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		err = database.RemovePath(record)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		lib.List(c, record.Path)
	})
}
