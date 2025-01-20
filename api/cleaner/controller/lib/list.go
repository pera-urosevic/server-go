package lib

import (
	"net/http"
	"server/api/cleaner/database"
	"server/api/cleaner/filesystem"
	"server/system"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context, path string) {
	entries, err := filesystem.Scan(path)
	if err != nil {
		system.GinError(c, err, true)
		return
	}

	records, err := database.GetPath(path)
	if err != nil {
		system.GinError(c, err, true)
	}

	for i, entry := range entries {
		for _, record := range records {
			if entry.Name == record.Name && path == record.Path {
				entries[i].OK = true
			}
		}
	}

	c.JSON(http.StatusOK, entries)
}
