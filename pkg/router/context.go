package router

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
)

func NewContext(address string) *Context {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	return &Context{
		Router:  r,
		Name:    "application",
		Address: address,
	}
}

func (c *Context) Start() error {
	log.Printf("starting server at %s", c.Address)
	return http.ListenAndServe(c.Address, c.Router)
}

func (c *Context) Get(path string, h Handler) {
	register(c, path, h, "GET")
}

func register(c *Context, path string, h Handler, method string) {
	c.Router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		err, response := h(c, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Errorf("request error %v", err)
			response = newError(err, r)
		}

		enc := json.NewEncoder(w)
		err = enc.Encode(response)
		if err != nil {
			w.Write([]byte("error getting json encoder"))
			return
		}

	}).Methods(method)

}

func newError(e error, r *http.Request) RouterError {
	return RouterError{
		Err:    e.Error(),
		URL:    r.RequestURI,
		Method: r.Method,
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requested %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
