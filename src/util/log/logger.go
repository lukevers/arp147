package log

import (
	L "log"
	"os"
)

var (
	Stdout = L.New(os.Stdout, "[arp147] ", L.LstdFlags)
	Stderr = L.New(os.Stderr, "[arp147] ", L.LstdFlags)
)
