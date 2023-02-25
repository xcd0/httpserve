package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./"))
	})
	http.HandleFunc("/aaa/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/aaa/\n")
	})
	if len(os.Args) == 1 {
		log.Fatal(http.ListenAndServe(":80", nil))
	} else if len(os.Args) == 2 {
		log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
	}
}
