package main

import (
	_ "arp147/ui/fonts"
	"flag"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	flag.Parse()

	if *debug {
		go http.ListenAndServe(":3030", http.DefaultServeMux)
	}
}
