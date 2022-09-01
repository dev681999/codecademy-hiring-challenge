package httpwriter

import (
	"bytes"
	"catinator-backend/pkg/model"
	"encoding/json"
	"net/http"
)

func WriteErrJsonResponse(statusCode int, w http.ResponseWriter, msg string) {
	WriteJsonResponse(statusCode, w, model.Error{
		Code:    statusCode,
		Message: msg,
	})
}

func Write200JsonResponse(w http.ResponseWriter, v any) {
	WriteJsonResponse(http.StatusOK, w, v)
}

func WriteJsonResponse(statusCode int, w http.ResponseWriter, v any) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(buf.Bytes())
}
