package storage

import (
	"os"
)

func AgendaDir() string {
	home, present := os.LookupEnv("HOME")
	if !present {
		home = "."
	}
	return home + "/.agenda/"
}

func SessionFile() string {
	return AgendaDir() + "session.json"
}

func CreateAgendaDir() {
	os.Mkdir(AgendaDir(), 0755)
}

func RemoveSessionFile() {
	err := os.Remove(SessionFile())
	// 若出错，则有可能是session.json本身是不存在的，也有可能是进程占用问题
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
}
