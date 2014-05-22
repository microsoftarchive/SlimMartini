package slim

import (
	"github.com/go-martini/martini"
)

type Handler struct {
	*martini.Martini
	martini.Router
}

func NewHandler() *Handler {
	r := martini.NewRouter()
	m := martini.New()
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &Handler{m, r}
}
