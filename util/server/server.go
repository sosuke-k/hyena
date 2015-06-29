package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/sosuke-k/hyena/util/git"
	"github.com/sosuke-k/hyena/util/pm"
)

func getHyenaPath() (hyenaPath string) {
	hyenaPath = os.Getenv("HYENA_DIR_PATH")
	if hyenaPath == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		hyenaPath = path.Join(usr.HomeDir, ".config/hyena")
	}
	return
}

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

	configPath := path.Join(getHyenaPath(), "config.json")

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

	projectDir := path.Join(getHyenaPath(), name)
	fmt.Println(projectDir)
	logString := git.Log(projectDir)

	if logString == "" {
		logString = "this project is not initialized"
	}

	w.Write([]byte(logString))
}

// Listen start web app
func Listen(port string) {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", homeHandler)
	rtr.HandleFunc("/project/{name}", projectHandler)

	http.Handle("/", rtr)
	http.ListenAndServe(":"+port, nil)
}
