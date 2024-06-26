package main

import (
  "fmt"
  "os"
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

  port := ":" + os.Getenv("PORT")

  if port == ":" {
    port += "8080"
  }

  log.Println(fmt.Sprintf("Serving on localhost%s", port))
  log.Fatal(http.ListenAndServeTLS(port, 
                                "https-certs/mzfaqiri.cert.pem",
                                "https-certs/private.key.pem", 
                                nil))
}

