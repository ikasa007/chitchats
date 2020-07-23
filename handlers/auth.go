package handlers

import (
	"fmt"
	"github.com/tim/chitchat/common/encrypt"
	"github.com/tim/chitchat/common/log"
	"github.com/tim/chitchat/models"
	"net/http"
)

// GET login
func Login(w http.ResponseWriter, r *http.Request){
	//t := ParseTemplateFiles("auth/layout","public/navbar","auth/login")
	//t.ExecuteTemplate(w,"layout",nil)
	GenerateHtml(w,nil,"auth/layout","public/navbar","auth/login")
}

func Signup(w http.ResponseWriter, r *http.Request ){
	GenerateHtml(w,nil,"auth/layout","public/navbar","auth/signup")
}

// register
func SignupAccount(w http.ResponseWriter , r *http.Request ){
	err := r.ParseForm()
	if err != nil {
		log.Crash("Can't Parse Form",err.Error())
	}
	Name := r.FormValue("name")
	Passwd := r.FormValue("password")
	Email := r.FormValue("email")

	user := models.User{
		Name:Name,
		Passwd:encrypt.Encrypt(Passwd),
		Email:Email,
	}

	if err := user.Create();err != nil {
		log.Crash(" Can't Create User",err.Error())
	}

	http.Redirect(w,r,"/login",http.StatusFound)

}

func Authenticate (w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Crash("Can't Parse Form",err.Error())
	}

	email := r.FormValue("email")
	passwd := r.FormValue("password")
	user,err := models.UserFindByEmail(email)
	if err != nil {
		log.Info("Can't find the User in the method ",err.Error())
	}
	if user.Passwd == encrypt.Encrypt(passwd) {
		// session
		session,err := user.CreateSession()
		if err != nil {
			log.Warning("Can't create session ")
		}
		cookie := http.Cookie{
			Name:"_cookie",
			Value:session.Uuid,
			HttpOnly:true,
		}
		http.SetCookie(w,&cookie)
		http.Redirect(w,r,"/",http.StatusFound)

	}else{
		http.Redirect(w,r,"/login",http.StatusFound)
	}

}

func Logout(w http.ResponseWriter , r *http.Request){
	cookie,err := r.Cookie("_cookie")
	fmt.Println(err)
	if err != http.ErrNoCookie {
		session := models.Session{Uuid:cookie.Value}
		session.DeleteByUuid()
	}
	http.Redirect(w,r,"/login",http.StatusFound)
}
