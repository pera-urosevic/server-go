package controller

import (
	"server/api/cleaner/controller/lib"
	"server/api/cleaner/database"
	"server/api/cleaner/types"
	"server/system"

	"github.com/gin-gonic/gin"
)

func RemovePath(r *gin.Engine) {
	r.DELETE("/cleaner", func(c *gin.Context) {
		var record = types.RecordCleaner{}
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
