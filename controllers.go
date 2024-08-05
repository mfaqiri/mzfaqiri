package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

type technology struct {
  Name string
  Summary string
  BulletPoints []string
}

var technologies = []technology {
  {
    Name: "git",
    Summary: "A version control system used in many projects to keep track of changes within the code, " +
            "collaborate with other developers, and sometimes revert changes entirely",
    BulletPoints: []string{
          "I use this technology constantly in my personal development.",
          "I have written scripts to automate some git actions",
          "3+ years of professional experience",
          "In many situations I have created a git repo, or gave recommendations in the creation of a git repo",
          "In a professional setting I make changes, and pull requests with git to present for a code review",
          "In a professional setting I have been asked to prepare branches for releasing a new version of our product"},
  },
  {
    Name: "C#",
    Summary: "An OOP language developed by Microsoft that is usually developed with a JIT (Just In Time) compiler" +
             ", to develop code in a variety of situations such as API/Web, IoT, Mobile/Desktop, Machine Learning," +
             "and Game development",
    BulletPoints: []string{
          "I have 3+ years of professional experience maintaining a deployed C# API application",
          "I have migrated a codebase from .NET Framework 4.6.2 -> 4.8 and then to .NET Core 6",
          "I have mentored colleagues on how to develop in our C# codebase",
          "I have worked on our deployment methods to bring it up to date with .NET 6"},
  },
  {
    Name: "bash",
    Summary: "A POSIX compliant shell used to interface with linux/unix environments",
    BulletPoints: []string{
          "I have used this extensively on my home linux machines",
          "Used to automate build processes professionally, and in personal projects",
          "Used to interface with remote machines using SSH, to do anything I may need to do remotely",
          "Experience configuring and setting environment variables"},
  },
  {
    Name: "Docker",
    Summary: "Docker is a platform for developing, shipping, and running applications. It is very useful as you do " +
          "not need to worry about the environment you are deploying to, docker will be setup with all the tooling " +
          "you have configured, and can test is working locally.",
    BulletPoints: []string{
          "I use docker in my personal projects, where I have to deploy to another environment",
          "I have experience setting up a docker node, with required environment details and setting up github actions to deploy when I push to my main branch"},
  },
}

var home_page_paragraph = "Hello, welcome to my website." +
                          "The purpose of this site is to " +
                          "present my development abilities " +
                          "to those interested in contacting " +
                          "me my email is linked below, as well " +
                          "as my github."

func main() {

  log.Println("app started")

  r := mux.NewRouter()

  fs := http.FileServer(http.Dir("assets/"))

  r.Handle("/", templ.Handler(home_page(home_page_paragraph)))
  r.Handle("/home", templ.Handler(home_page(home_page_paragraph)))
  r.Handle("/resume", templ.Handler(resume_page()))
  r.Handle("/certifications", templ.Handler(certification_page()))
  r.Handle("/technology", templ.Handler(technology_page(technologies, 0)))
  r.HandleFunc("/technology/{techInd}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    techInd,err := strconv.Atoi(vars["techInd"])

    if err != nil {
      log.Fatal(err)

    }  else {

        log.Print(technologies[techInd])

        technology_page_summary(technologies[techInd]).Render(r.Context(), w)

      }

  })

  r.Handle("/static/js/htmx.min.js", http.StripPrefix("/static",fs))
  r.Handle("/static/css/skeleton.css", http.StripPrefix("/static",fs))
  r.Handle("/static/images/email.svg", http.StripPrefix("/static",fs))
  r.Handle("/static/images/github.svg", http.StripPrefix("/static",fs))
  r.Handle("/static/images/apple-touch-icon.png", http.StripPrefix("/static",fs))
  r.Handle("/static/images/favicon-32x32.png", http.StripPrefix("/static",fs))
  r.Handle("/static/images/favicon-16x16.png", http.StripPrefix("/static",fs))
  r.Handle("/static/images/site.webmanifest", http.StripPrefix("/static",fs))

  r.Handle("/calendar", templ.Handler(calendar_page()))
  r.Handle("/about-me", templ.Handler(about_me_page()))

  port := ":" + os.Getenv("PORT")

  if port == ":" {
    port += "8080"
  }

  log.Println(fmt.Sprintf("Serving on localhost%s", port))
  log.Fatal(http.ListenAndServeTLS(port,
                                "https-certs/mzfaqiri.cert.pem",
                                "https-certs/private.key.pem",
                                r))
}

