package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"server/api/blog/database"
	"server/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdatePost(r *gin.Engine) {
	r.PUT("/blog/post/:postID", func(c *gin.Context) {
		postID, err := strconv.ParseInt(c.Param("postID"), 10, 64)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		type updateBody struct {
			URL         string `json:"url"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Image       string `json:"image"`
			Category    string `json:"category"`
			Template    string `json:"template"`
		}
		payload := updateBody{}
		err = json.Unmarshal(body, &payload)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		post, err := database.GetPost(postID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		post.URL = payload.URL
		post.Title = payload.Title
		post.Description = payload.Description
		post.Image = payload.Image
		post.Category = payload.Category
		post.Template = payload.Template

		err = database.UpdatePost(post)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
