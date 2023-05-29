package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func init(){
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthHandler)
}

func main(){
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request){
	status := 200
	w.WriteHeader(status)
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	addr := r.RemoteAddr
	ip := strings.Split(addr, ":")[0]
	fmt.Printf("root %s %d\n", ip, status)
}


func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}