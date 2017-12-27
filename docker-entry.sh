#!/bin/bash
usage() {
    echo "docker run agenda-service (server | client | repl) args..."
}

case $1 in
    server)
        shift
        /go/bin/server "$@"
        ;;
    client)
        shift
        /go/bin/client "$@"
        ;;
    repl)
        exec /bin/bash
        ;;
    *)
        usage
        ;;
esac