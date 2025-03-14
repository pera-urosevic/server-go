package controller

import (
	"server/api/cleaner/controller/lib"
	"server/api/cleaner/database"
	"server/api/cleaner/database/model"
	"server/system"

	"github.com/gin-gonic/gin"
)

func AddPath(r *gin.Engine) {
	r.PUT("/cleaner", func(c *gin.Context) {
		var record = model.Cleaner{}
		err := c.BindJSON(&record)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		err = database.AddPath(record)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		lib.List(c, record.Path)
	})
}
