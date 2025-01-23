package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"server/api/gallery/database"
	"server/api/gallery/database/model"
	"server/api/gallery/meta"
	"server/api/gallery/rename"
	"server/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdatePhoto(r *gin.Engine) {
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

		var photoNew model.Photo
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
}
