package handlerfuncs

import (
	"encoding/json"
	"net/http"

	"github.com/yashdalfthegray/hue-remote-v2/responses"
)

// Status sends out the status of the server as a health check
func Status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(responses.Status{
		Status: "ok",
	}); err != nil {
		panic(err)
	}
}
