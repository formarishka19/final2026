package api

import (
	"encoding/json"
	"errors"
	"final2026/pkg/db"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	ID string `json:"id"`
}
type ResponseError struct {
	Error string `json:"error"`
}

func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err, ok := data.(error); ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseError{Error: err.Error()})
	} else {
		json.NewEncoder(w).Encode(Response{ID: fmt.Sprint(data)})
	}

}

func checkDate(task *db.Task) error {
	now := time.Now()
	if len(task.Date) == 0 {
		task.Date = now.Format(DATEFORMAT)
	}
	t, err := time.Parse(DATEFORMAT, task.Date)
	if err != nil {
		return err
	}
	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		return err
	}
	if AfterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(DATEFORMAT)
		} else {
			task.Date = next
		}
	}
	return nil
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t db.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		writeJson(w, err)
		return
	}
	if len(t.Title) == 0 {
		writeJson(w, errors.New("Title is empty"))
		return
	}
	err = checkDate(&t)
	if err != nil {
		writeJson(w, err)
		return
	}
	taskId, err := db.AddTask(&t)
	if err != nil {
		writeJson(w, err)
		return
	}
	writeJson(w, taskId)
}
