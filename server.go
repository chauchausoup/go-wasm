package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8888", "listen address")
	dir    = flag.String("dir", ".", "directory to ser")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))))
}
