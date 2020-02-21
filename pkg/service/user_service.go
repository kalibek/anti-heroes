package service

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kalibek/anti-heroes/pkg/repo"
	"github.com/kalibek/anti-heroes/pkg/router"
)

func SaveUser(r repo.UserRepo) router.Handler {
	return func(req *http.Request) (interface{}, error) {
		v := mux.Vars(req)
		id, err := strconv.ParseInt(v[id], 10, 64)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return r.GetUser(id)
	}
}
