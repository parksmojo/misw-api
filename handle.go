package main

import (
	"misw/middleware"
	"net/http"
)

func Handle(route string, handler http.HandlerFunc) {
	http.Handle(route, middleware.LogRequest(handler))
}