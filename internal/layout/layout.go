package layout

import (
	"net/http"

	"github.com/a-h/templ"
)

func Handler(content templ.Component, options ...func(*templ.ComponentHandler)) http.Handler {
	return templ.Handler(BaseLayout(content), options...)
}
