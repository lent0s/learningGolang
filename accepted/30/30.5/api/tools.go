package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// convert Request Body from json to *Api
func readRequest(r *http.Request, req *Api) error {

	content, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	if err = json.Unmarshal(content, &req); err != nil {
		return err
	}
	return nil
}

// check for not supported method
func notSupportedMethod(request string, support string, w http.ResponseWriter) bool {
	if request != support {
		w.Header().Set("Support", support)
		http.Error(w, "method not allowed", 405)
		return true
	}
	return false
}
