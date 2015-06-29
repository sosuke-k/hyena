package server

import (
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/sosuke-k/hyena/util/pm"
)

// Page strcut
type Page struct {
	Projects []string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	root := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/root")
	templatePath := path.Join(root, "index.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	hyenaPath := os.Getenv("HYENA_DIR_PATH")
	if hyenaPath == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		hyenaPath = path.Join(usr.HomeDir, ".config/hyena")
	}
	configPath := path.Join(hyenaPath, "config.json")

	projects := pm.Load(configPath)
	page := Page{Projects: projects}
	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Here is " + name + " project page."))
}

// Listen start web app
func Listen() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", homeHandler)
	rtr.HandleFunc("/project/{name:[a-z]+}", projectHandler)

	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}
