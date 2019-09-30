package api

import (
	"fmt"
	"net/http"
)

type responseFunc func(http.ResponseWriter, *http.Request) error
type inFunc func(http.ResponseWriter, *http.Request)

func AddHeaders(f responseFunc, contentType string) inFunc {
	return func(w http.ResponseWriter, r *http.Request) { // the return must match the full DoPrintSomething function signature
		w.Header().Set("Content-Type", contentType)
		err := f(w, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}
