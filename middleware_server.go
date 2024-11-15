package main

import (
	"fmt"
	"net/http"
	"time"
)

func middlewareServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v: %s %s\n", time.Now().UTC().Format(time.DateTime), r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
