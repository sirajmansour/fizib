package main

import (
	"fmt"
	"log"
	"net/http"
	"website"
)

func main() {
	website.Port = ":8081"
	srv := website.NewServer()
	http.HandleFunc("/welcome/", srv.Welcome())
	http.HandleFunc("/", srv.SignUp())
	http.HandleFunc("/login/", srv.LogIn())

	fmt.Printf("Muxless server is up.\n")
	if err := http.ListenAndServe(website.Port, nil); err != nil {
		log.Fatal(err)
	}
}
