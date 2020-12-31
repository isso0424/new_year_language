package handler

import (
	"bytes"
	"fmt"
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
	fmt.Println(result)
	w.Write([]byte(result))
}
