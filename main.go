package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	f, err := os.OpenFile("public/index.html", os.O_CREATE|os.O_RDWR, os.ModePerm)

	defer f.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Listening http://localhost:8000")

	http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}