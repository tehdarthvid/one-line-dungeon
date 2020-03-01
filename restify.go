package restify

import (
	"fmt"
	"net/http"
	// "darthvid.io/old/apis/encounter"
)

// HTTPGetEncounter is an HTTP Cloud Function intgrator.
func HTTPGetEncounter(w http.ResponseWriter, r *http.Request) {
	strJSON := "yo!"
	if "" != strJSON {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprint(w, strJSON)
	} else {
		http.Error(w, "iamerror", 500)
	}
}
