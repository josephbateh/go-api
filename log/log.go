package log

import "log"

// Info logs with an INFO prefix
func Info(value string) {
	log.Println("- INFO -", value)
}

// Error logs with an ERROR prefix
func Error(value string) {
	log.Println("- ERROR -", value)
}

// Fatal logs with an ERROR prefix
func Fatal(value string) {
	log.Fatalln("- Fatal -", value)
}
