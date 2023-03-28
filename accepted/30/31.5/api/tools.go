package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// convert Request Body from json to *Api
func readRequest(r *http.Request, req *Api) error {

	content, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	req.list = make(map[string]interface{})
	if err = json.Unmarshal(content, &req.list); err != nil {
		return err
	}
	return checkCorrectReq(req)
}

// check for correct request
func checkCorrectReq(req *Api) error {

	var fields = []string{"name", "age", "source_id", "target_id", "new age"}

	for field, val := range req.list {
		contain := false
		for _, key := range fields {
			if field == key {
				contain = true
				break
			}
		}
		if !contain {
			return fmt.Errorf("incorrect field [%q], allowed: %q",
				field, fields)
		}

		switch val.(type) {
		case string:
			if field != fields[0] {
				return fmt.Errorf("incorrect data in field [%q]", field)
			}
		default:
			if _, err := strconv.Atoi(fmt.Sprintf("%v", val)); err != nil {
				return err
			}
		}
	}
	return nil
}

// check for not supported method
func notSupportedMethod(request string, support string,
	w http.ResponseWriter) bool {

	if request != support {
		w.Header().Set("Support", support)
		err := fmt.Sprintf("method \"%s\" not allowed [%s]", request, support)
		http.Error(w, err, http.StatusMethodNotAllowed)
		return true
	}
	return false
}
