package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	var err error
	if repeat == "" {
		return "", nil
	}
	date, err := time.Parse(DATEFORMAT, dstart)
	if err != nil {
		return "", errors.New("error of parsing dstart to time value")
	}
	params := strings.Split(repeat, " ")
	years := 0
	days := 0
	switch params[0] {
	case "d":

		if len(params) < 2 {
			return "", errors.New("invalid format of repeat rule")
		}
		days, err = strconv.Atoi(params[1])
		if err != nil {
			return "", errors.New("invalid quanity of days in rule")
		}
		if days > 400 {
			return "", errors.New("invalid quanity of days in rule, nust be less than 400")
		}
	case "y":
		years = 1
	default:
		return "", errors.New("invalid format of rule")
	}
	for {
		date = date.AddDate(years, 0, days)
		if AfterNow(date, now) {
			break
		}
	}
	return date.Format(DATEFORMAT), nil
}

func AfterNow(date, now time.Time) bool {
	nowStr := now.Format(DATEFORMAT)
	dateStr := date.Format(DATEFORMAT)
	if dateStr > nowStr {
		return true
	}
	return false
}

func NextDateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "only GET requests allowed", http.StatusBadRequest)
		return
	}

	now, err := time.Parse(DATEFORMAT, r.FormValue("now"))
	if err != nil {
		now = time.Now()
	}
	nDate, err := NextDate(now, r.FormValue("date"), r.FormValue("repeat"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err != nil {
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nDate))

}
