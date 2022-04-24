package main

import (
	"fmt"
	"log"
	"net/http"
	// "wxcloudrun-golang/db"
	// "wxcloudrun-golang/service"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func main() {
	// if err := db.Init(); err != nil {
	// 	panic(fmt.Sprintf("mysql init failed with %+v", err))
	// }

	http.HandleFunc("/", handler)
	// http.HandleFunc("/api/count", service.CounterHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
