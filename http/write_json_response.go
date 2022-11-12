package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, content interface{}, status int) {
	b, err := json.Marshal(content)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	io.WriteString(w, string(b))
}
