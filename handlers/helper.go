package handlers

import (
	"errors"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	. "github.com/tim/chitchat/config"
	"github.com/tim/chitchat/models"
	"html/template"
	"net/http"
	"strings"
	"time"
)
//var config *Configuration
var Localizer *i18n.Localizer
func init(){
	// load localizer
	//config = LoadConfig()
	Localizer = i18n.NewLocalizer(ViperConfig.LocaleBundle,ViperConfig.App.Lang)


}

func Session(w http.ResponseWriter , r *http.Request) (session models.Session,err error){
	cookie, err := r.Cookie("_cookie")
	fmt.Println("Cookie:",cookie)
	if err == nil {
		session = models.Session{Uuid:cookie.Value}
		if ok,_ := session.Check(); !ok {
			err = errors.New(" Invalid Session ")
		}
	}
	return
}

func ParseTemplateFiles(filenames ...string ) (t *template.Template ){
	var files []string
	for _,file := range filenames {
		files = append(files,fmt.Sprintf("views/%s/%s.html",ViperConfig.App.Lang,file))
	}
	t = template.Must(template.ParseFiles(files ...))
	return

}

func GenerateHtml (w http.ResponseWriter, data interface{},filenames ...string) {
	var files []string
	for _,file := range filenames {
		files = append(files,fmt.Sprintf("views/%s/%s.html",ViperConfig.App.Lang,file))
	}
	fmt.Println(files)
	funcMap := template.FuncMap{"fdate":FormatDate}
	t := template.New("layout").Funcs(funcMap)
	templates := template.Must(t.ParseFiles(files ...))
	templates.ExecuteTemplate(w,"layout",data)

}

func Version() string {
	return "v1.0"
}

func FormatDate(t time.Time) string {
	dateTime := "2006-01-02 15:04:05"
	return t.Format(dateTime)
}

func ErrMsg(w http.ResponseWriter , r *http.Request,msg string){
	url := []string{"/err?msg=",msg}
	http.Redirect(w,r,strings.Join(url,""),http.StatusFound)
}


