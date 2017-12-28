# Agenda-Service

![travis](https://travis-ci.org/MarshallW906/Agenda-Service.svg?branch=dev)

## Introdution
Agenda-Service for Golang-Server-Computing

Group Members:
- [FideoJ](https://github.com/FideoJ)
- [ZM-J](https://github.com/ZM-J)
- [MarshallW906](https://github.com/MarshallW906)

## Setup
### Host
```
go get -v github.com/MarshallW906/Agenda-Service/server github.com/MarshallW906/Agenda-Service/client \
  && go install github.com/MarshallW906/Agenda-Service/server \
  && go install github.com/MarshallW906/Agenda-Service/client
$GOPATH/bin/server &
go test -v ./server/vendor/model/
go test -v ./client/vendor/service/
$GOPATH/bin/client -h
```

### Docker
```
docker pull fideo/agenda-service
docker run -d -p 8080:8080 -v /server_data:/data --rm --name agenda_server fideo/agenda-service server
alias agenda_cli='docker run --net host -v /cli_data:/data --rm fideo/agenda-service client'
agenda_cli -h
```
