package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/schema"
)

var (
	// ErrInvalidContentType is returned by ParseRequest if it can't unmarshal it into the passed struct
	ErrInvalidContentType = errors.New("Invalid request content type")

	// gorilla/schema decoder is a shared object, as it caches information about structs
	decoder = schema.NewDecoder()
)

// JSONResponse writes JSON to an http.ResponseWriter with the corresponding status code
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	// Get rid of the invalid status codes
	if status < 100 || status > 599 {
		status = 200
	}

	// Try to marshal the input
	result, err := json.Marshal(data)
	if err != nil {
		// Set the result to the default value to prevent empty responses
		result = []byte(fmt.Sprintf(`{"status":500,"message":"Error occured while marshalling the response body : %s"}`, err))
	}

	// Set the response's content type to JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Write the result
	w.WriteHeader(status)
	w.Write(result)
}

func parseIP(ipStr string) net.IP {
	ip := net.ParseIP(ipStr)
	return ip
}

// ParseRequest takes the input body from the passed request and tries to unmarshal it into data
func ParseRequest(r *http.Request, data interface{}) error {
	// Get the contentType for comparsions
	contentType := r.Header.Get("Content-Type")

	// Deterimine the passed ContentType
	if strings.Contains(contentType, "application/json") {
		// It's JSON, so read the body into a variable
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		// And then unmarshal it into the passed interface
		err = json.Unmarshal(body, data)
		if err != nil {
			return err
		}

		return nil
	} else if contentType == "" ||
		strings.Contains(contentType, "application/x-www-form-urlencoded") ||
		strings.Contains(contentType, "multipart/form-data") {
		// net/http should be capable of parsing the form data
		err := r.ParseForm()
		if err != nil {
			return err
		}

		// Unmarshal them into the passed interface
		err = decoder.Decode(data, r.PostForm)
		if err != nil {
			return err
		}

		return nil
	}

	return ErrInvalidContentType
}
