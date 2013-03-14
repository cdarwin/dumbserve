package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var (
	webroot = flag.String("webroot", "/tmp", "absolute path to main directory of files to serve")
	altroot = flag.String("altroot", "", "absolute path to alternate directory of files to serve")
	port    = flag.String("port", "8080", "port to run server on")
)

func updateRoot(s string) interface{} {
	str := "cd " + s + " && git pull origin master"
	out, err := exec.Command("/bin/sh", "-c", str).Output()
	if err != nil {
		log.Println(err)
	}
	return out
}

func github(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", updateRoot(*webroot))
	fmt.Fprintf(w, "%s", updateRoot(*altroot))
}

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*webroot)))
	http.Handle("/alt/", http.StripPrefix("/alt/", http.FileServer(http.Dir(*altroot))))
	http.HandleFunc("/github", github)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Println(err)
	}
}
