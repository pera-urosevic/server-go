package gallery

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"server/api/gallery/database"
	"server/api/gallery/flick"
	"server/api/gallery/meta"
	"server/api/gallery/rename"
	"server/api/gallery/sync"
	"server/api/gallery/types"
	"server/system"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/gallery/", func(c *gin.Context) {
		photos, err := database.GetPhotos("")
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, photos)
	})

	r.GET("/gallery/:filter", func(c *gin.Context) {
		filter := c.Param("filter")
		photos, err := database.GetPhotos(filter)
		if err != nil {
			system.GinError(c, err, false)
			return
		}
		c.JSON(http.StatusOK, photos)
	})

	r.GET("/gallery/sync", func(c *gin.Context) {
		photos, err := sync.Sync()
		if err != nil {
			system.GinError(c, err, false)
			return
		}
		c.JSON(http.StatusOK, photos)
	})

	r.GET("/gallery/deploy", func(c *gin.Context) {
		deploy := os.Getenv("GALLERY_DEPLOY")
		cmd := strings.Split(deploy, "|")
		app := exec.Command(cmd[0], cmd[1:]...)
		err := app.Start()
		app.Process.Release()
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	r.GET("/gallery/open/:photoPath", func(c *gin.Context) {
		photoPath := c.Param("photoPath")
		app := exec.Command("explorer.exe", photoPath)
		err := app.Start()
		app.Process.Release()
		c.JSON(http.StatusOK, err)
	})

	r.GET("/gallery/photo/:photoID", func(c *gin.Context) {
		photoID, err := strconv.ParseInt(c.Param("photoID"), 10, 64)
		if err != nil {
			system.GinError(c, err, false)
			return
		}
		post, err := database.GetPhoto(photoID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, post)
	})

	r.GET("/gallery/photo/:photoID/flickr", func(c *gin.Context) {
		photoID, err := strconv.ParseInt(c.Param("photoID"), 10, 64)
		if err != nil {
			system.GinError(c, err, false)
			return
		}
		photo, err := database.GetPhoto(photoID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		res, err := flick.Upload(photo)
		if err != nil {
			system.GinError(c, err, true)
		} else {
			c.JSON(http.StatusOK, res)
		}
	})

	r.PUT("/gallery/photo/:photoID", func(c *gin.Context) {
		photoID, err := strconv.ParseInt(c.Param("photoID"), 10, 64)
		if err != nil {
			system.GinError(c, err, false)
			return
		}
		photoOld, err := database.GetPhoto(photoID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		var photoNew types.Photo
		err = json.Unmarshal(body, &photoNew)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		if photoNew.ID != photoOld.ID {
			system.GinError(c, errors.New("photo ID mismatch"), false)
			return
		}
		photoNew.Path, err = rename.Rename(photoOld, photoNew)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		photoNew.Modified, err = meta.Update(photoNew)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		err = database.UpdatePhoto(photoNew)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, photoNew)
	})

	return r
}
