package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/sosuke-k/hyena/util/git"
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

func projectListAPIHandler(w http.ResponseWriter, r *http.Request) {
	projects := pm.Projects{}
	projectList := pm.Load(path.Join(getHyenaPath(), "config.json"))

	for _, name := range projectList {
		projects = append(projects, pm.Project{Name: name})
	}

	json.NewEncoder(os.Stdout).Encode(projects)

	if err := json.NewEncoder(w).Encode(projects); err != nil {
		panic(err)
	}
}