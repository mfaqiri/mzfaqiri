package main

import (
	"log"
	"net/http"
	"os"
    "fmt"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

type technology struct {
  Name string
  Summary string
  Usage string
}

var technologies = []technology {
  {
    Name: "GIT",
    Summary: "A version control system used in many projects to keep track of changes within the code, " +
            "collaborate with other developers, and sometimes revert changes entirely",

    Usage: "This is my default choice for a VCS (Version Control System), I use this for whatever code base I am in even the one for this very website.",
  },
  {
    Name: "C#",
    Summary: "An OOP language developed by Microsoft that is usually developed with a JIT (Just In Time) compiler" +
             ", to develop code in a variety of situations such as API/Web, IoT, Mobile/Desktop, Machine Learning," +
             "and Game development",

    Usage : "I primarily would use C# for API development, something that I feel needs a more robust codebase to maintain and add features",
  },
  {
    Name: "BASH",
    Summary: "A POSIX compliant shell used to interface with linux/unix environments",

    Usage: "I use this shell to navigate my personal machines, make changes in the file system. Sometimes I may make a script to automate a large command I use often.",
  },
  {
    Name: "DOCKER",
    Summary: "Docker is a platform for developing, shipping, and running applications. It is very useful as you do " +
          "not need to worry about the environment you are deploying to, docker will be setup with all the tooling " +
          "you have configured, and can test is working locally.",

    Usage: "I primarily use Docker to create applications where I would like to have a virtual environment without much access to files outside of what the application needs. This web application is built and deployed with Docker.",
  },
  {
    Name: "GOLANG",
    Summary: "An open source programming language created by GOOGLE",

    Usage: "This website is the first application I have built with go, and it is a language I have taken a great liking to. It has a robust set of features built with it, and with the help of templ I have a more integrated experience using html with the benefits of golang",
  },
}

var home_page_paragraph = "Hello, my name is Mansoor Faqiri and I am a Full Stack Developer."+
                            " I am aiming to become an entry level CyberSecurity professional, I am CompTIA Security + certified,"+
                            " proof of this is in the certifications tab in the form of an embedded badge."+
                            " Other than this certification I am also personally studying various cybersecurity concepts through HackTheBox."+
                            " If you are interested in my professional experience as a developer, you may check my resume tab for a downloadable PDF of my resume, with work experience."+
                            " To see when I am available, you can check my calendar for my availability."+
                            " To see how to message me details are on my resume and below are links to my github and a mailto link to my email,"+
                            " which will automatically open an email to send to me if you have a program lik e.g. Outlook, Thunderbird, etc. installed"+
                            " Thank you for visiting my site"

func main() {

  log.Println("app started")

  r := mux.NewRouter()

  fs := http.FileServer(http.Dir("assets/"))

  r.Handle("/", templ.Handler(home_page(home_page_paragraph)))
  r.Handle("/home", templ.Handler(home_page(home_page_paragraph)))
  r.Handle("/resume", templ.Handler(resume_page()))
  r.Handle("/certifications", templ.Handler(certification_page()))
  r.Handle("/technology", templ.Handler(technology_page(technologies)))
  r.Handle("/static/js/htmx.min.js", http.StripPrefix("/static",fs))
  r.Handle("/static/css/skeleton.css", http.StripPrefix("/static",fs))
  r.Handle("/static/images/email.svg", http.StripPrefix("/static",fs))
  r.Handle("/static/images/github.svg", http.StripPrefix("/static",fs))
  r.Handle("/static/images/linkedin.svg", http.StripPrefix("/static",fs))
  r.Handle("/static/images/apple-touch-icon.png", http.StripPrefix("/static",fs))
  r.Handle("/static/images/favicon-32x32.png", http.StripPrefix("/static",fs))
  r.Handle("/static/images/favicon-16x16.png", http.StripPrefix("/static",fs))
  r.Handle("/static/images/site.webmanifest", http.StripPrefix("/static",fs))
  r.Handle("/static/images/professional_portrait.jpg", http.StripPrefix("/static",fs))

  r.Handle("/calendar", templ.Handler(calendar_page()))

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

