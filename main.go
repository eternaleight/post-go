package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type dogdogdog int

func (m dogdogdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	log.Println(req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	_ = exec.Command("open", "http://localhost:8080").Start()
	var d dogdogdog
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, d)
}
