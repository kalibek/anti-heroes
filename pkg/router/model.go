package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Context struct {
	Router  *mux.Router
	Name    string
	Address string
}

type Handler func(c *Context, r *http.Request) (error, interface{})

type RouterError struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	Err    string `json:"error"`
}
