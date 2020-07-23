package main

import (
	"github.com/tim/chitchat/common/log"
	. "github.com/tim/chitchat/config"
	"github.com/tim/chitchat/routes"
	"net/http"
	"strconv"
)

func main(){
	Start()
}

func Start(){
	// load config
//	config := LoadConfig()
	r := routes.NewRouters()

	// static file requests
	assets := http.FileServer(http.Dir(ViperConfig.App.Static))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",assets))

	http.Handle("/",r)
	log.Info("Start the server on port "+ViperConfig.App.Addr + strconv.Itoa(ViperConfig.App.Port))
	err := http.ListenAndServe(":" + strconv.Itoa(ViperConfig.App.Port),nil)

	if err != nil {
		log.Info("An error occured starting HTTP listener at port " + strconv.Itoa(ViperConfig.App.Port))
		log.Crash("Error: " + err.Error())
	}

}

