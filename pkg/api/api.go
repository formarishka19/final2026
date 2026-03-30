package api

import "net/http"

const DATEFORMAT = "20060102"

func taskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPost:
		AddTaskHandler(w, r)
	case http.MethodDelete:
		return
	default:
		http.Error(w, "only GET/POST requests allowed", http.StatusBadRequest)
	}

	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(nDate))

}

func Init(mux *http.ServeMux) {
	mux.HandleFunc("/api/nextdate", NextDateHandler)
	mux.HandleFunc("/api/task", taskHandler)
}
