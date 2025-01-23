package image

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"server/api/blog/database/model"
	"server/api/blog/log"

	"github.com/disintegration/imaging"
)

func pathImages() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s\\Data\\Blog\\Images", home)
}

func pathOriginal(post model.Post) string {
	return fmt.Sprintf("%s\\original\\%s-image", pathImages(), post.Timestamp)
}

func pathThumbnail(post model.Post) string {
	return fmt.Sprintf("%s\\thumbnail\\%s-image.jpg", pathImages(), post.Timestamp)
}

func pathExcerpt(post model.Post) string {
	return fmt.Sprintf("%s\\excerpt\\%s-image.jpg", pathImages(), post.Timestamp)
}

func pathPost(post model.Post) string {
	return fmt.Sprintf("%s\\post\\%s-image.jpg", pathImages(), post.Timestamp)
}

func Fetch(post model.Post) error {
	res, err := http.Get(post.Image)
	if err != nil {
		log.Log(err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Log(err)
		return err
	}

	originalPath := pathOriginal(post)

	file, err := os.Create(originalPath)
	if err != nil {
		log.Log(err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Log(err)
		return err
	}

	imageOriginal, err := imaging.Open(originalPath)
	if err != nil {
		log.Log(err)
		return err
	}

	postPath := pathPost(post)
	imagePost := imaging.Fit(imageOriginal, 640, 640, imaging.Lanczos)
	err = imaging.Save(imagePost, postPath, imaging.JPEGQuality(75))
	if err != nil {
		log.Log(err)
	}

	excerptPath := pathExcerpt(post)
	imageExcerpt := imaging.Fit(imageOriginal, 320, 320, imaging.Lanczos)
	err = imaging.Save(imageExcerpt, excerptPath, imaging.JPEGQuality(75))
	if err != nil {
		log.Log(err)
	}

	thumbnailPath := pathThumbnail(post)
	imageThumbnail := imaging.Fill(imageOriginal, 160, 140, imaging.Center, imaging.Lanczos)
	err = imaging.Save(imageThumbnail, thumbnailPath, imaging.JPEGQuality(75))
	if err != nil {
		log.Log(err)
	}

	return nil
}

func Delete(post model.Post) error {
	err := os.Remove(pathOriginal(post))
	if err != nil {
		log.Log(err)
	}

	err = os.Remove(pathThumbnail(post))
	if err != nil {
		log.Log(err)
	}

	err = os.Remove(pathExcerpt(post))
	if err != nil {
		log.Log(err)
	}

	err = os.Remove(pathPost(post))
	if err != nil {
		log.Log(err)
	}

	return nil
}
