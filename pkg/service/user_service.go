package service

import (
	"github.com/gorilla/mux"
	"github.com/kalibek/anti-heroes/pkg/repo"
	"github.com/kalibek/anti-heroes/pkg/router"
	"net/http"
)

func SaveUser(r *repo.Repo) router.Handler {
	return func (req *http.Request) (interface{}, error) {
		mux.Vars(req)
		r.GetUser()
	}
}
