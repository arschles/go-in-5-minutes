package handlers

import (
	"fmt"
	"net/http"
)

func jsonErr(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
}
