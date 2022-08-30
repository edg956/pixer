package controllers

import (
	"fmt"
	"net/http"
)

func CreateNewAlbum(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}
