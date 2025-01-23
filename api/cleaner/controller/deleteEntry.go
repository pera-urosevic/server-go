package controller

import (
	"server/api/cleaner/controller/lib"
	"server/api/cleaner/database/model"
	"server/api/cleaner/filesystem"
	"server/system"

	"github.com/gin-gonic/gin"
)

func DeleteEntry(r *gin.Engine) {
	r.DELETE("/cleaner/delete", func(c *gin.Context) {
		var record = model.Cleaner{}
		err := c.BindJSON(&record)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		err = filesystem.Delete(record)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		lib.List(c, record.Path)
	})
}
