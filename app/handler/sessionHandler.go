package handler

import (
	"GoogleProjectBackend/util"
	"fmt"
	"net/http"
)

func SessionAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SessionAuthMiddleware")
		session, err := util.Store.Get(r, "survaySession")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		userId := session.Values["authenticated"]

		if session.IsNew || userId != "true" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
