package sync

import (
	"encoding/json"
	"os/exec"
	"reflect"
	"server/api/gallery/log"
	"server/api/gallery/types"
	"strings"
	"time"
)

func metaTitle(exif map[string]interface{}) string {
	title, ok := exif["ImageDescription"]
	if ok {
		return title.(string)
	}
	return ""
}

func metaDatetime(exif map[string]interface{}, file types.AlbumFile) string {
	datetime, ok := exif["DateTimeOriginal"]

	if ok && datetime != nil && datetime != "" {
		temp := datetime.(string)
		temp = strings.ReplaceAll(temp, ":", "")
		temp = strings.ReplaceAll(temp, " ", "")
		return temp
	}

	modifiedTime := time.UnixMilli(file.Modified)
	return modifiedTime.Format("20060102150405")
}

func metaDescription(exif map[string]interface{}) string {
	description, ok := exif["UserComment"]

	if ok {
		return description.(string)
	}

	return ""
}

func metaKeywords(exif map[string]interface{}) string {
	keywords, ok := exif["Keywords"]

	if ok {

		if reflect.TypeOf(keywords).Kind() == reflect.String {
			temp := strings.Split(keywords.(string), ",")
			return strings.Join(temp, " | ")
		}

		if reflect.TypeOf(keywords).Kind() == reflect.Slice {
			keywordInterfaces := keywords.([]interface{})
			keywordStrings := []string{}
			for _, keywordInterface := range keywordInterfaces {
				keywordStrings = append(keywordStrings, keywordInterface.(string))
			}

			return strings.Join(keywordStrings, " | ")
		}
	}

	return ""
}

func metaCopyright(exif map[string]interface{}) string {
	copyright, ok := exif["Copyright"]

	if ok {
		return copyright.(string)
	}

	return ""
}

func readMeta(file types.AlbumFile) (types.Photo, error) {
	run := exec.Command("exiftool.exe", "-json", file.Path)
	output, err := run.CombinedOutput()
	if err != nil {
		log.Log(err)
		return types.Photo{}, err
	}

	results := []map[string]interface{}{}
	json.Unmarshal(output, &results)
	exif := results[0]

	record := types.Photo{
		Path:        file.Path,
		Album:       file.Album,
		Modified:    file.Modified,
		Online:      file.Online,
		Type:        file.Path[len(file.Path)-3:],
		Title:       metaTitle(exif),
		Datetime:    metaDatetime(exif, file),
		Description: metaDescription(exif),
		Keywords:    metaKeywords(exif),
		Copyright:   metaCopyright(exif),
	}

	return record, nil
}
