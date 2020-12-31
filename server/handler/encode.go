package handler

import (
	"bytes"
	"net/http"

	"new_year_language/translater"
)

func EncodeHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	text := buf.String()
	result := translater.Translate(text, false)
	w.Write([]byte(result))
}
