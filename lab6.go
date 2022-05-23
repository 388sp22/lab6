package lab6

import (
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	// TODO: echo back original bytes of body if the value of `command` in the body is `echo`.
	// Because r.FormValue reads the body, this might be a bit tricky.
}
