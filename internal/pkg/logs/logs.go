package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func Initiliaze(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(f)
	return nil
}
