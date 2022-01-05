package http

import "net/http"

// Init function with all Endpoints off application
func Init() {
	http.ListenAndServe(":9000", nil)
}
