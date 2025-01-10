package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"server/api/blog/database"
	"server/system"

	"github.com/gin-gonic/gin"
)

func CreatePost(r *gin.Engine) {
	r.POST("/blog/post", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		type createBody struct {
			URL string `json:"url"`
		}
		payload := createBody{}
		err = json.Unmarshal(body, &payload)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		id, err := database.InsertPost(payload.URL)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		type createResponse struct {
			ID int64 `json:"id"`
		}
		c.JSON(http.StatusOK, createResponse{ID: id})
	})
}
