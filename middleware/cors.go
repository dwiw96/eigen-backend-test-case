package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Cors(next httprouter.Handle) httprouter.Handle {
	return (func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.Method == "OPTIONS" {
			next(w, r, ps)
			return
		}

		next(w, r, ps)
	})
}
