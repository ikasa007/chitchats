package config

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"os"
	"sync"
)

type App struct {
	Addr string
	Static string
	Log string
	Port int
	Locale string
	Lang string

}

type Database struct {
	Driver string
	Addr string
	Port int
	Database string
	User string
	Passwd string
	ParseTime bool
	Charset string
}
// why not use pointer
type Configuration struct {
	App App
	Db  Database
	LocaleBundle *i18n.Bundle
}

var config *Configuration
var once sync.Once

//load configuration
func LoadConfig() *Configuration{
	once.Do(func() {
		file,err := os.Open("config.json")
		if err != nil {
			log.Fatalln("Can't load config file",err)
		}
		decode := json.NewDecoder(file)
		fmt.Println(decode)
		config = &Configuration{}
		err = decode.Decode(config)
		if err != nil {
			log.Fatalln("Can't get config from file",err)
		}

		// locale loading
		bundle := i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("json",json.Unmarshal)
		bundle.LoadMessageFile(config.App.Locale+"/active.en.json")
		bundle.MustLoadMessageFile(config.App.Locale+"/active."+config.App.Lang+".json")
		config.LocaleBundle = bundle

	})

	return config
}
