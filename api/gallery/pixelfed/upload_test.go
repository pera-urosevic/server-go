package pixelfed

import (
	"server/api/gallery/database"
	"testing"
)

func TestUpload(t *testing.T) {
	photo, err := database.GetPhoto(1)
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}
	res, err := Upload(photo)
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}
	if len(res) == 0 {
		t.Errorf("res is empty")
		return
	}
}
