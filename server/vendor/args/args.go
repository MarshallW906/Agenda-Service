package args

import (
	"github.com/spf13/pflag"
)

// Port : Server listening on target Port
var Port *string

// Host : Server listening on target Host
var Host *string

// DBfile for agenda
var DBfile *string

func init() {
	Port = pflag.StringP("port", "p", "8080", "Port of the server listening on")
	Host = pflag.StringP("host", "h", "0.0.0.0", "Host of the server listening on")
	DBfile = pflag.StringP("db", "d", "agenda.db", "Path to the sqlite db file")
	pflag.Parse()
}
