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
	// http.HandleFunc("/", srv.SignUp())
	http.HandleFunc("/login/", srv.LogIn())

	http.Handle("/static/",
		http.StripPrefix("/static",
			http.FileServer(
				http.Dir(`C://HtmlPages/MUXLESS/signUp`))))

	fmt.Printf("Muxless server is up.\n")
	if err := http.ListenAndServe(website.Port, nil); err != nil {
		log.Fatal(err)
	}
}
