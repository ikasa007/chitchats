package log

func Crash(args ...interface{}){
	logger.SetPrefix("CRASH ")
	logger.Println(args...)
}

func Info(args ...interface{}){
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func Warning (args ...interface{}){
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}
