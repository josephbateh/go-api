package verbs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Get is used for an HTTP GET request
func Get(response http.ResponseWriter, request *http.Request, v interface{}) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	response.Header().Set("Content-Type", "application/json")

	// Check if the method is a get
	if request.Method != http.MethodGet {
		http.Error(response, http.StatusText(405), 405)
		fmt.Println(response)
		return
	}

	enc := json.NewEncoder(response)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
}

// Post is used for an HTTP POST request
func Post(response http.ResponseWriter, request *http.Request, v interface{}) {
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	response.Header().Set("Content-Type", "application/json")

	// Check if the method is a get
	if request.Method != http.MethodPost {
		http.Error(response, http.StatusText(405), 405)
		fmt.Println(response)
		return
	}

	enc := json.NewEncoder(response)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
}

// Put is used for an HTTP PUT request
func Put(response http.ResponseWriter, request *http.Request, v interface{}) {
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	response.Header().Set("Content-Type", "application/json")

	// Check if the method is a get
	if request.Method != http.MethodPut {
		http.Error(response, http.StatusText(405), 405)
		fmt.Println(response)
		return
	}

	enc := json.NewEncoder(response)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
}

// Patch is used for an HTTP PATCH request
func Patch(response http.ResponseWriter, request *http.Request, v interface{}) {
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "PATCH,OPTIONS")
	response.Header().Set("Content-Type", "application/json")

	// Check if the method is a get
	if request.Method != http.MethodPatch {
		http.Error(response, http.StatusText(405), 405)
		fmt.Println(response)
		return
	}

	enc := json.NewEncoder(response)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
}

// Delete is used for an HTTP DELETE request
func Delete(response http.ResponseWriter, request *http.Request, v interface{}) {
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "DELETE,OPTIONS")
	response.Header().Set("Content-Type", "application/json")

	// Check if the method is a get
	if request.Method != http.MethodDelete {
		http.Error(response, http.StatusText(405), 405)
		fmt.Println(response)
		return
	}

	enc := json.NewEncoder(response)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
}
