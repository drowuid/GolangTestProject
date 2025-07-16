package handlers 

import (

"encoding/json"
"net/http"
"time"
)

type Status struct {

	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
	TimeZone string `json:"timezone"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	loc := time.Now().Location()
	status := Status{
		Message: "Server is running smoothly",
		Timestamp: time.Now().Format(time.RFC3339),
		TimeZone: loc.String(),
}

json.NewEncoder(w).Encode(status)
}
