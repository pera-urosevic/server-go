package blog

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"server/api/blog/database"
	"server/api/blog/scraper"
	"server/api/blog/scraper/image"
	"server/system"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/blog/", func(c *gin.Context) {
		posts, err := database.GetPosts("")
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	r.GET("/blog/:filter", func(c *gin.Context) {
		filter := c.Param("filter")
		posts, err := database.GetPosts(filter)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	r.GET("/blog/deploy", func(c *gin.Context) {
		deploy := os.Getenv("BLOG_DEPLOY")
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

	r.GET("/blog/post/:postID", func(c *gin.Context) {
		postID, err := strconv.ParseInt(c.Param("postID"), 10, 64)
		if err != nil {
			system.GinError(c, err, false)
			return
		}
		post, err := database.GetPost(postID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, post)
	})

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

	r.PATCH("/blog/post/:postID", func(c *gin.Context) {
		postID, err := strconv.ParseInt(c.Param("postID"), 10, 64)
		if err != nil {
			system.GinError(c, err, false)
			return
		}
		post, err := database.GetPost(postID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		post, err = scraper.Scrape(post)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		err = database.UpdatePost(post)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, post)
	})

	r.DELETE("/blog/post/:postID", func(c *gin.Context) {
		postID, err := strconv.ParseInt(c.Param("postID"), 10, 64)
		if err != nil {
			system.GinError(c, err, false)
			return
		}
		post, err := database.GetPost(postID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		err = image.Delete(post)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		err = database.RemovePost(postID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	return r
}
