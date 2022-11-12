package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody[T any](w http.ResponseWriter, r *http.Request, t *T) error {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return err
	}

	// Unmarshal
	err = json.Unmarshal(b, &t)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return err
	}

	return nil
}
