package storage

import (
	"os"

	"entity"
	"logger"
)

func LoadCurToken() (string, bool) {
	file, err := os.OpenFile(SessionFile(), os.O_RDONLY|os.O_CREATE, 0644)
	logger.FatalIf(err)
	session := entity.DeserializeSession(file)

	return session.CurrentToken, session.CurrentToken != ""
}

func StoreSession(session *entity.Session) {
	file, err := os.OpenFile(SessionFile(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	logger.FatalIf(err)
	session.Serialize(file)
}
