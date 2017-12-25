package server

import (
	"args"

	"github.com/codegangsta/negroni"
)

// Start the server
func Start() {
	newServer().Run(*args.Host + ":" + *args.Port)
}

// NewServer configures and returns a Server.
func newServer() *negroni.Negroni {
	n := negroni.Classic()
	n.UseHandler(router())
	return n
}
