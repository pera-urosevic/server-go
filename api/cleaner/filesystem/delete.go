package filesystem

import (
	"os"
	"os/exec"
	"server/api/cleaner/database/model"
	"server/api/cleaner/log"
	"strings"
)

func Delete(record model.Cleaner) error {
	root := os.Getenv("CLEANER_ROOT")
	path := root + record.Path + "/" + record.Name
	path = strings.ReplaceAll(path, "/", "\\")

	cmd := "recycle-bin.exe"
	args := []string{path}

	log.Log("[DELETE]", path)

	run := exec.Command(cmd, args...)
	_, err := run.CombinedOutput()
	if err != nil {
		log.Log(err)
	}
	return err
}
