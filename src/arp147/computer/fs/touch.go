package fs

import (
	"time"
)

func Touch(file string) {
	vfs[file].UpdatedAt = time.Now()
}
