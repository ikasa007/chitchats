package posts

import (
	"fmt"
	"github.com/tim/chitchat/common/log"
	"github.com/tim/chitchat/handlers"
	"github.com/tim/chitchat/models"
	"net/http"
)

func PostThread(w http.ResponseWriter , r *http.Request) {
	sess, err := handlers.Session(w,r)
	if err != nil {
		http.Redirect(w,r,"/",http.StatusFound)
	}else{
		err := r.ParseForm()
		if err != nil {
			log.Crash("Can not parse form")
		}
		user,err := sess.SessionUser()
		if err != nil {
			log.Crash("Can not get user from session in method post thread")
		}

		body := r.FormValue("body")
		uuid := r.FormValue("uuid")
		thread,err := models.ThreadByUuid(uuid)
		if err != nil {
			log.Crash("Can not get thread through uuid ")
		}

		if _,err := user.CreatePosts(thread,body); err != nil {
			log.Crash("Can not create post in controller")
			handlers.ErrMsg(w,r,"Can not create post in controller")
		}
		url := fmt.Sprint("/thread/read?id=",uuid)
		http.Redirect(w,r,url,http.StatusFound)

	}


}
