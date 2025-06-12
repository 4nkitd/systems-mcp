package log

import "log"

func Write(level string, message string) {

	log.Printf("[%s] %s\n", level, message)
}
