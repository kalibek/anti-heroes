package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Context struct {
	Router  *mux.Router
	Address string
}

type Handler func(r *http.Request) (interface{}, error)

type Error struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	Err    string `json:"error"`
}
