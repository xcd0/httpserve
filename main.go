package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile) // ログの出力書式を設定する
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./"))
	})
	http.HandleFunc("/aaa/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/aaa/\n")
	})
	log.Fatal(http.ListenAndServe(":80", nil)) // リバースプロキシサーバーの起動
}
