package fileserve

import (
	"catinator-backend/pkg/httpwriter"
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

var noCacheHeaders = map[string]string{
	"Cache-Control": "no-cache, private, max-age=0",
	"Pragma":        "no-cache",
}

func NewHandler(urlParam string, folder string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileName := chi.URLParam(r, urlParam)
		file := filepath.Join(folder, fileName)
		exists, err := fileExists(file)
		if err != nil {
			httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
			return
		}

		if !exists {
			httpwriter.WriteErrJsonResponse(http.StatusNotFound, w, "Request file not found")
			return
		}

		// Set our NoCache headers
		for k, v := range noCacheHeaders {
			w.Header().Set(k, v)
		}
		http.ServeFile(w, r, file)
	}
}

func fileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
