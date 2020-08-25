package http200

import (
	"fmt"
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "http200")
	fmt.Fprintf(w, "ok")
}
