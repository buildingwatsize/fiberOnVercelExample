package handler

import (
	"fmt"
	"net/http"
)

func HTTPPackageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{"msg":"Hello World!"}`)
}
