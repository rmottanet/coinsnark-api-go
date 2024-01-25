package middleware

import "net/http"

type ConvertHandlerFunc func(http.ResponseWriter, *http.Request)


func ConvertHandler(fn ConvertHandlerFunc) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fn(w, r)
    })
}
