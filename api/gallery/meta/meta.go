package meta

import (
	"fmt"
	"os"
	"os/exec"
	"server/api/gallery/database/model"
	"strings"
)

func title(params []string, title string) []string {
	params = append(params, fmt.Sprintf("-EXIF:ImageDescription=%s", title))
	params = append(params, fmt.Sprintf("-IPTC:Headline=%s", title))
	return params
}

func description(params []string, description string) []string {
	params = append(params, fmt.Sprintf("-EXIF:UserComment=%s", description))
	params = append(params, fmt.Sprintf("-IPTC:Caption-Abstract=%s", description))
	return params
}

func keywords(params []string, keywords string) []string {
	for _, keyword := range strings.Split(keywords, " | ") {
		params = append(params, fmt.Sprintf("-IPTC:Keywords=%s", keyword))
	}
	return params
}

func datetime(params []string, datetime string) []string {
	dt := datetime[0:4] + ":" + datetime[4:6] + ":" + datetime[6:8] + " " + datetime[8:10] + ":" + datetime[10:12] + ":" + datetime[12:14]
	// 2019:12:21 11:34:10
	params = append(params, fmt.Sprintf("-EXIF:DateTimeOriginal=%s", dt))
	// 2019:12:21
	params = append(params, fmt.Sprintf("-IPTC:DateCreated=%s", dt[0:10]))
	// 11:34:10+00:00
	params = append(params, fmt.Sprintf("-IPTC:TimeCreated=%s+00:00", dt[11:19]))
	return params
}

func copyright(params []string, copyright string) []string {
	params = append(params, fmt.Sprintf("-EXIF:Copyright=%s", copyright))
	params = append(params, fmt.Sprintf("-IPTC:CopyrightNotice=%s", copyright))
	return params
}

func pixelfed(params []string, pixelfed string) []string {
	params = append(params, fmt.Sprintf("-EXIF:Pixelfed=%s", pixelfed))
	return params
}

func Update(photo model.Photo) (int64, error) {
	params := []string{"-m", "-overwrite_original", "-XMP=", "-Orientation=", "-gps:all="}
	params = title(params, photo.Title)
	params = description(params, photo.Description)
	params = keywords(params, photo.Keywords)
	params = datetime(params, photo.Datetime)
	params = copyright(params, photo.Copyright)
	params = pixelfed(params, photo.Pixelfed)
	params = append(params, photo.Path)

	run := exec.Command("exiftool.exe", params...)
	_, err := run.CombinedOutput()
	if err != nil {
		return 0, err
	}

	fileInfo, err := os.Stat(photo.Path)
	if err != nil {
		return 0, err
	}

	return fileInfo.ModTime().UnixMilli(), nil
}
