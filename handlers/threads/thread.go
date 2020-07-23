package threads

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/tim/chitchat/common/log"
	"github.com/tim/chitchat/handlers"
	"github.com/tim/chitchat/models"
	"net/http"
)

//

func CreateThreads(w http.ResponseWriter , r *http.Request){
	_,err := handlers.Session(w,r)

	// show the template
	if err != nil {
		http.Redirect(w,r,"/login",http.StatusFound)
	}else{
		handlers.GenerateHtml(w,nil,"home/layout","auth/navbar","threads/create")
	}


}

func PostsThreads(w http.ResponseWriter , r *http.Request){
	sess,err := handlers.Session(w,r)
	if err != nil {
		http.Redirect(w,r,"/login",http.StatusFound)
	}else{
		err = r.ParseForm()
		if err != nil {
			log.Crash("Can't parse Form ",err.Error())
		}
		user,err := sess.SessionUser()
		if err != nil {
			log.Crash("Can't got the user message ",err)
		}
		topic := r.PostFormValue("topic")

		_,err = user.CreateThread(topic)
		if err != nil {
			log.Crash("Can't Create thread",err)
		}
		http.Redirect(w,r,"/",http.StatusFound)

	}

}


func GetThreads(w http.ResponseWriter,r *http.Request){
	vals := r.URL.Query()
	uuid := vals.Get("id")
	thread,err := models.ThreadByUuid(uuid)
	if err != nil {
		log.Crash("can't got thread")
		msg := handlers.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID:"thread_not_found",
		})
		handlers.ErrMsg(w, r, msg)
	}else{
		_,err := handlers.Session(w,r)
		if err != nil {
			fmt.Println("unlogin")
			handlers.GenerateHtml(w,&thread,"home/layout","public/navbar","threads/index")
		}else{
			fmt.Println("login")
			handlers.GenerateHtml(w,&thread,"home/layout","auth/navbar","auth/thread")

		}

	}



}


