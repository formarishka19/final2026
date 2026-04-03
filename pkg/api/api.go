package api

import "net/http"

const DATEFORMAT = "20060102"

func taskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		AddTaskHandler(w, r)
	case http.MethodGet:
		ViewTaskHandler(w, r)
	case http.MethodPut:
		UpdateTaskHandler(w, r)
	case http.MethodDelete:
		DeleteTaskHandler(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusBadRequest)
	}

}
func taskListkHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		ListTaskHandler(w, r)
	default:
		http.Error(w, "only GET requests allowed", http.StatusBadRequest)
	}

}

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		DoneTaskHandler(w, r)
	default:
		http.Error(w, "only POST requests allowed", http.StatusBadRequest)
	}

}

func Init(mux *http.ServeMux) {
	mux.HandleFunc("/api/nextdate", NextDateHandler)
	mux.HandleFunc("/api/task", taskHandler)
	mux.HandleFunc("/api/tasks", taskListkHandler)
	mux.HandleFunc("/api/task/done", taskDoneHandler)
}
