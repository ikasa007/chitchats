package log

import (
	"log"
	"os"
)

var logger *log.Logger


// todo config the log file and split as the date
func init(){
	file,err := os.OpenFile("logs/chitchat.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	if err != nil {
		log.Fatalln("Can not open log file",err)
	}
	logger = log.New(file,"INFO ",log.Ldate|log.Ltime|log.Lshortfile,)
}
