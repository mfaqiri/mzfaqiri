package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
)


func main() {

  log.Println("app started")



  http.Handle("/", templ.Handler(resume_page()))
  http.Handle("/home", templ.Handler(home_page()))
  http.Handle("/resume", templ.Handler(resume_page()))
  http.Handle("/technology", templ.Handler(technology_page()))
  http.Handle("/calendar", templ.Handler(calendar_page()))
  http.Handle("/about-me", templ.Handler(about_me_page()))

  log.Fatal(http.ListenAndServe(":8000", nil))
}

