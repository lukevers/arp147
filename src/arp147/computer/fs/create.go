package fs

import (
	"errors"
	"util/log"
)

func Create(path string, file Item) error {
	if Exists(path) {
		return errors.New("File already exists")
	}

	vfs[path] = &file

	log.Stdout.Println(vfs)

	return nil
}
