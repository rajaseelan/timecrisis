package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/rajaseelan/timecrisis/pkg"
)

func home(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	machineHostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Something went wrong trying to get the Hostname", 500)
		return
	}

	tStruct := pkg.TimeStruct{
		String: t.Format(time.UnixDate),
		Unix:   t.Unix(),
		Host:   machineHostname,
	}

	bytes, err := json.MarshalIndent(&tStruct, "", "  ")
	if err != nil {
		http.Error(w, "Unable to Marshal Response", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}
