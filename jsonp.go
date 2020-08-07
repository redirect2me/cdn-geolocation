package main

import (
	"encoding/json"
	"net/http"
)

func write_with_callback(w http.ResponseWriter, r *http.Request, v interface{}) {

	callback := r.URL.Query().Get("callback")

	var b []byte
	var err error
	b, err = json.MarshalIndent(v, "", "  ")
	if err != nil {
		b = []byte("{\"success\":false,\"message\":\"json.Marshal failed\"}")
	}

	if callback > "" {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(callback))
		w.Write([]byte("("))
		w.Write(b)
		w.Write([]byte(");"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		w.Header().Set("Access-Control-Max-Age", "604800") // 1 week
		w.Write(b)
	}
}
