package auth

import (
	// "fmt"
	// "html/template"
	"log"
	"net/http"
)

func RegisterRoutes(handler *Handler) {
	log.Println("Routes registered")
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//fmt.Println(r.Method)
			// temp,err:=template.ParseFiles("./web/template/register.html")
			// if err!=nil{
			// 	fmt.Println("error:",err)
			// 	return
			// }
			http.ServeFile(w, r, "web/template/register.html")
			// temp.Execute(w,nil)
			return
		}
		handler.Register(w, r)
	
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "web/template/login.html")
			return
		}
		handler.Login(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		http.ServeFile(w, r, "web/template/home.html")
	})
}
