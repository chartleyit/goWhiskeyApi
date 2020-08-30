package p

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json: "message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello HTTP!")
		return
	}
	if d.Message == "" {
		fmt.Fprint(w, "Hello HTTP!")
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Message))
}
