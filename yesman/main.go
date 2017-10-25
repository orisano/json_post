package main

import (
	"flag"
	"log"
	"net"
	"net/http"
)

type yesman struct{}

func (*yesman) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	var path string
	flag.StringVar(&path, "p", "yesman.sock", "unix domain socket path")
	flag.Parse()

	d, err := net.Listen("unix", path)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	log.Fatal(http.Serve(d, &yesman{}))
}
