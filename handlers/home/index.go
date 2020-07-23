package home

import (
	"github.com/tim/chitchat/handlers"
	"github.com/tim/chitchat/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	// got all threads from database and render to the template
	threads, err := models.Thread()
	// todo add some debug
	if err == nil {
		_,err := handlers.Session(w,r)
		if err != nil {
			handlers.GenerateHtml(w,threads,"home/layout","public/navbar","home/index")
		}else{
			handlers.GenerateHtml(w,threads,"home/layout","auth/navbar","home/index")
		}
	}

}


func Err(w http.ResponseWriter, r *http.Request){
	vals := r.URL.Query()

	_,err := handlers.Session(w,r)
	if err != nil {
		handlers.GenerateHtml(w,vals.Get("msg"),"home/layout","public/navbar","public/error")
	}else{
		handlers.GenerateHtml(w,vals.Get("msg"),"auth/layout","public/navbar","public/error")

	}

}
