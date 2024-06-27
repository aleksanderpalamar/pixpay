package logger

import (
	"log"
	"os"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
