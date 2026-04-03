package api

import (
	"encoding/json"
	"final2026/pkg/db"
	"net/http"
	"time"
)

type TaskListResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func ListTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(TaskListResp{
		Tasks: tasks,
	})
}

func ViewTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := r.FormValue("id")
	if taskID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "no ID in request"})
		return
	}
	t := db.Task{ID: taskID}
	err := db.GetTask(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(t)
}

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := r.FormValue("id")
	if taskID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "no ID in request"})
		return
	}
	t := db.Task{ID: taskID}
	err := db.GetTask(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}
	if t.Repeat == "" {
		err := db.DeleteTask(&t)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Error: err.Error()})
			return
		}

	} else {
		now := time.Now()
		nextDate, err := NextDate(now, t.Date, t.Repeat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Error: "error nextdate calculating"})
			return
		}

		t.Date = nextDate
		err = db.UpdateTask(&t)

	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(Response{})
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := r.FormValue("id")
	if taskID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "no ID in request"})
		return
	}
	t := db.Task{ID: taskID}
	err := db.DeleteTask(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(Response{})
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t db.Task
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}
	if len(t.Title) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: `empty title`})
		return
	}
	err = checkDate(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}
	err = db.UpdateTask(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(Response{})
}
