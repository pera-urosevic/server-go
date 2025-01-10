package controller

import (
	"server/api/cleaner/controller/lib"
	"server/api/cleaner/filesystem"
	"server/api/cleaner/types"
	"server/system"

	"github.com/gin-gonic/gin"
)

func DeleteEntry(r *gin.Engine) {
	r.DELETE("/cleaner/delete", func(c *gin.Context) {
		var record = types.RecordCleaner{}
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
