package main

import (
	"github.com/brankad/api"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"time"
)

func serveHomepage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writingSync.Lock()
	programIsRunning = true
	writingSync.Unlock()

	var homepage api.HomePage
	homepage.Time = time.Now().String()


	tmpl := template.Must(template.ParseFiles("ui/html/homepage.html"))
	_ = tmpl.Execute(writer, homepage)

	writingSync.Lock()
	programIsRunning = false
	writingSync.Unlock()
}
