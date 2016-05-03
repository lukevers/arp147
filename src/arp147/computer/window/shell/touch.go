package shell

import (
	"arp147/computer/fs"
	"arp147/computer/line"
	"errors"
	"strings"
	"time"
)

type Touch struct {
	Shell
}

func (t *Touch) Name() string {
	return "touch"
}

func (t *Touch) Description() string {
	return "Change file access and modification times"
}

func (t *Touch) Run(command *line.Command) error {
	if len(command.Arguments) < 1 {
		return errors.New("Not enough arguments")
	}

	// Touch each file
	for _, file := range command.Arguments {
		// Get absolute path if it's not an absolute path already.
		if !strings.HasPrefix(file, "/") {
			trailing := strings.HasSuffix(command.Directory, "/")
			if trailing {
				file = command.Directory + file
			} else {
				file = command.Directory + "/" + file
			}
		}

		// Split up to find all directories
		parts := strings.Split(file, "/")
		last := len(parts) - 1

		for i, _ := range parts {
			// Figure out the full path of this part of the path.
			f := ""
			if i > 0 {
				for j := 1; j <= i; j++ {
					f += "/" + parts[j]
				}
			} else {
				f = "/"
			}

			if fs.Exists(f) {
				if i == last {
					// Change time to now
					fs.Touch(f)
					break
				} else {
					continue
				}
			} else {
				if i == last {
					// Make sure we aren't trying to create a directory from
					// using the touch utility.
					if strings.HasSuffix(f, "/") {
						return errors.New("No such directory")
					}

					// Create file
					now := time.Now()
					return fs.Create(f, fs.Item{
						Contents:  "",
						Dir:       false,
						User:      command.User,
						Group:     command.Group,
						Mode:      644,
						CreatedAt: now,
						UpdatedAt: now,
					})
				} else {
					return errors.New("No such directory")
				}
			}
		}
	}

	return nil
}
