package cleaner

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"server/api/cleaner/database"
	"server/api/cleaner/filesystem"
	"server/api/cleaner/types"
	"server/system"
	"strings"

	"github.com/gin-gonic/gin"
)

func list(c *gin.Context, path string) {
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
			if entry.Name == record.Name {
				entries[i].OK = true
			}
		}
	}

	c.JSON(http.StatusOK, entries)
}

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/cleaner/bookmarks", func(c *gin.Context) {
		bookmarks, err := database.GetBookmarks()
		if err != nil {
			system.GinError(c, err, true)
		}

		c.JSON(http.StatusOK, bookmarks)
	})

	r.POST("/cleaner/bookmarks", func(c *gin.Context) {
		var bookmark = ""
		err := c.BindJSON(&bookmark)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		err = database.AddBookmark(bookmark)
		if err != nil {
			system.GinError(c, err, true)
		}

		bookmarks, err := database.GetBookmarks()
		if err != nil {
			system.GinError(c, err, true)
		}

		c.JSON(http.StatusOK, bookmarks)
	})

	r.DELETE("/cleaner/bookmarks", func(c *gin.Context) {
		var bookmark = ""
		err := c.BindJSON(&bookmark)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		err = database.RemoveBookmark(bookmark)
		if err != nil {
			system.GinError(c, err, true)
		}

		bookmarks, err := database.GetBookmarks()
		if err != nil {
			system.GinError(c, err, true)
		}

		c.JSON(http.StatusOK, bookmarks)
	})

	r.GET("/cleaner/open/*path", func(c *gin.Context) {
		path := c.Param("path")[1:]
		root := os.Getenv("CLEANER_ROOT")
		path = root + path
		path = strings.ReplaceAll(path, "/", "\\")

		cmd := os.Getenv("CLEANER_OPEN")
		if cmd == "" {
			err := fmt.Errorf("CLEANER_OPEN not set")
			system.GinError(c, err, true)
			return
		}

		app := exec.Command(cmd, path)
		err := app.Start()
		if err == nil {
			app.Process.Release()
			c.JSON(http.StatusOK, nil)
			return
		}

		system.GinError(c, err, true)
	})

	r.GET("/cleaner/list/*path", func(c *gin.Context) {
		path := c.Param("path")[1:]
		list(c, path)
	})

	r.PUT("/cleaner", func(c *gin.Context) {
		var record = types.RecordCleaner{}
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

		list(c, record.Path)
	})

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

		list(c, record.Path)
	})

	return r
}
