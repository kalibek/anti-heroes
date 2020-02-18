package main

import (
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/kalibek/anti-heroes/pkg/router"
)

func main() {
	c := router.NewContext(":9000")
	c.Get("/", hello)
	c.Get("/err", err)

	log.Fatal(c.Start())
}

func hello(c *router.Context, r *http.Request) (error, interface{}) {
	return nil, fmt.Sprintf("method: %s, app: %s", r.Method, c.Name)
}

func err(c *router.Context, r *http.Request) (error, interface{}) {
	return errors.New("Something bad happened"), nil
}
