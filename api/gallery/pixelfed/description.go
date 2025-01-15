package pixelfed

import (
	"fmt"
	"server/api/gallery/types"
	"strings"
)

func getDescription(photo types.Photo) string {
	line1 := photo.Title

	tags := strings.Split(photo.Keywords, " | ")
	for i, tag := range tags {
		tags[i] = "#" + tag
	}
	line2 := strings.Join(tags, " ")

	description := fmt.Sprintf("%s\n%s", line1, line2)
	return description
}
