package fs

import (
	"time"
)

var vfs map[string]*Item

type Item struct {
	Contents  string
	Dir       bool
	User      string
	Group     string
	Mode      uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	now := time.Now()
	vfs = map[string]*Item{
		"/": &Item{
			Contents:  "",
			Dir:       true,
			User:      "root",
			Group:     "root",
			Mode:      777,
			CreatedAt: now,
			UpdatedAt: now,
		},
		"/README": &Item{
			Contents:  "Welcomeeeeee!",
			Dir:       false,
			User:      "root",
			Group:     "root",
			Mode:      777,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}
