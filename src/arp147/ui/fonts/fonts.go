package fonts

import (
	"github.com/golang/freetype/truetype"
	"io/ioutil"
	"path/filepath"
	"util/log"
)

var fonts map[string]*truetype.Font = make(map[string]*truetype.Font)

// When this package is loaded for the first time, it loads all packages inside
// of our font directory into an in-memory map for quicker usage. Since this is
// a function that is necessary, if anything produces an error the application
// will be shut down and the error will be logged to Stderr.
func init() {
	files, err := filepath.Glob("resources/fonts/*.ttf")
	if err != nil {
		log.Stderr.Fatal("Could not load font files", err)
	}

	for _, f := range files {
		_, name := filepath.Split(f)
		data, err := ioutil.ReadFile(f)
		if err != nil {
			log.Stderr.Fatal("Could not read font file", err)
		}

		fonts[name], err = truetype.Parse(data)
		if err != nil {
			log.Stderr.Fatal("Could not parse file as ttf", err)
		}
	}
}

func Get(font Font) *truetype.Font {
	return fonts[string(font)]
}
