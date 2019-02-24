package filesystem

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/blang/vfs/memfs"
)

type VirtualFS struct {
	cwd string
	FS  *memfs.MemFS
}

func New() *VirtualFS {
	fs := &VirtualFS{
		cwd: "/",
		FS:  memfs.Create(),
	}

	return fs.initialize()
}

func (fs *VirtualFS) initialize() *VirtualFS {
	log.Println("Generating filesystem")

	// TODO: normalize for windows, or pre-build binary with internal files.
	prefix := "terminal/filesystem/root"

	// TODO: permissions? (that aren't 777)

	err := filepath.Walk(
		prefix,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			vpath := strings.TrimPrefix(path, prefix)
			if info.IsDir() {
				if vpath != "" {
					fs.FS.Mkdir(vpath, 0777)
					log.Println("Creating:", vpath)
				}
			} else {
				file, err := fs.FS.OpenFile(vpath, os.O_CREATE|os.O_RDWR, 0777)
				if err != nil {
					return err
				}

				bytes, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}

				_, err = file.Write(bytes)
				if err != nil {
					return err
				}

				file.Close()
				log.Println("Creating:", vpath)
			}

			return nil
		},
	)

	if err != nil {
		log.Println(err)
	}

	return fs
}
