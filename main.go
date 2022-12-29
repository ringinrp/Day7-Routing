package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/Public/").Handler(http.StripPrefix("/Public", http.FileServer(http.Dir("./Public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/project", project).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/project-detail", projectDetail).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("POST")

	port := "8000"

	fmt.Println("Server sedang berjalan di port " + port)
	http.ListenAndServe("localhost:"+port, route)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-type", "text/html; charset-utf-8")

	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message: " + err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}
func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-type", "text/html; charset-utf-8")

	tmpl, err := template.ParseFiles("views/project.html")

	if err != nil {
		w.Write([]byte("Message: " + err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-type", "text/html; charset-utf-8")

	tmpl, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message: " + err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}
func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-type", "text/html; charset-utf-8")

	tmpl, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		w.Write([]byte("Message: " + err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}
func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Project Name : " + r.PostForm.Get("project")) //Get data from input name tag
	fmt.Println("Start Date : " + r.PostForm.Get("date-start"))
	fmt.Println("End Date : " + r.PostForm.Get("date-end"))
	fmt.Println("Description : " + r.PostForm.Get("description"))
	fmt.Println("Use Node JS : " + r.PostForm.Get("nodejs"))
	fmt.Println("Use React JS : " + r.PostForm.Get("reactjs"))
	fmt.Println("Use Next JS : " + r.PostForm.Get("nextjs"))
	fmt.Println("Use Vue JS : " + r.PostForm.Get("vuejs"))

	http.Redirect(w, r, "/project", http.StatusMovedPermanently)
}
