package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLog(path string) error{
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
	return nil
}