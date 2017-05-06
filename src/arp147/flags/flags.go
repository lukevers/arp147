package flags

import (
	"flag"
)

// Debug is a flag that turns on debugging.
var Debug = flag.Bool("debug", false, "Debug mode")

func init() {
	flag.Parse()
}
