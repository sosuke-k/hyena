package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"text/template"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"github.com/sosuke-k/hyena/util/git"
	"github.com/sosuke-k/hyena/util/log"
	"github.com/sosuke-k/hyena/util/pm"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templateDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/root/templates")
	templatePath := path.Join(templateDir, "index.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	projects := pm.Projects{}
	projectList := pm.Load(path.Join(getHyenaPath(), "config.json"))
	for _, name := range projectList {
		projects = append(projects, pm.Project{Name: name})
	}

	err = tmpl.Execute(w, projects)
	if err != nil {
		panic(err)
	}
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	hyenaLogger := logger.GetInstance()
	methodURL := r.Method + " " + r.URL.String()
	hyenaLogger.Println(methodURL)
	fmt.Fprintln(os.Stdout, methodURL)

	// params := mux.Vars(r)
	// name := params["name"]
	// projectDir := path.Join(getHyenaPath(), name)
	// shas := gitinfo.GetSHAArray(projectDir)
	// var commits []gitinfo.Commit
	// for _, sha := range shas {
	// 	commits = append(commits, *gitinfo.NewCommit(git.Show(projectDir, sha)))
	// }
	// fh := gitinfo.FileHistories{}
	// gitinfo.ConvertCommitsToHistory(commits, &fh)

	templateDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/root/templates")
	templatePath := path.Join(templateDir, "project.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		hyenaLogger.Fatalln(err)
		fmt.Fprintln(os.Stderr, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		hyenaLogger.Fatalln(err)
		fmt.Fprintln(os.Stderr, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hyenaLogger.Println("response wrote")
	fmt.Fprintln(os.Stdout, "response wrote")
}

func projectListAPIHandler(w http.ResponseWriter, r *http.Request) {
	projects := pm.Projects{}
	projectList := pm.Load(path.Join(getHyenaPath(), "config.json"))

	for _, name := range projectList {
		projects = append(projects, pm.Project{Name: name})
	}

	if err := json.NewEncoder(w).Encode(projects); err != nil {
		panic(err)
	}
}

// Result struct
type Result struct {
	Success bool `json:"success"`
}

func projectDeleteAPIHandler(w http.ResponseWriter, r *http.Request) {
	hyenaLogger := logger.GetInstance()
	methodURL := r.Method + " " + r.URL.String()
	hyenaLogger.Println(methodURL)
	fmt.Fprintln(os.Stdout, methodURL)

	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		hyenaLogger.Println("failed to convert body to json")
		hyenaLogger.Fatalln(err)
		fmt.Fprintln(os.Stderr, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	configPath := path.Join(getHyenaPath(), "config.json")
	for _, name := range js.Get("checked_list").MustArray() {
		pm.Delete(configPath, name.(string))
	}
	fmt.Fprintln(os.Stdout, "completed to delete")

	resultJSON, err := json.Marshal(Result{Success: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJSON)
	hyenaLogger.Println("response write " + string(resultJSON[:]))
	fmt.Fprintln(os.Stdout, "response write "+string(resultJSON[:]))
}

func projectDiffAPIHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	newCommit := params["newCommit"]
	oldCommit := params["oldCommit"]
	projectDir := path.Join(getHyenaPath(), name)
	rep := gyena.Repository{Dir: projectDir}
	diffString := rep.Diff(oldCommit, newCommit)

	if diffString == "" {
		diffString = "this project is not initialized or these commit not exist"
	}

	w.Write([]byte(diffString))
}

func projectHistoryAPIHandler(w http.ResponseWriter, r *http.Request) {
	hyenaLogger := logger.GetInstance()
	methodURL := r.Method + " " + r.URL.String()
	hyenaLogger.Println(methodURL)
	fmt.Fprintln(os.Stdout, methodURL)

	params := mux.Vars(r)
	name := params["name"]
	projectDir := path.Join(getHyenaPath(), name)
	rep := gyena.Repository{Dir: projectDir}
	shas := rep.GetSHAArray()
	var commits []gyena.Commit
	for _, sha := range shas {
		commits = append(commits, *gyena.NewCommit(rep.Show(sha)))
	}
	fh := gyena.FileHistories{}
	gyena.ConvertCommitsToHistory(commits, &fh)

	if err := json.NewEncoder(w).Encode(fh); err != nil {
		hyenaLogger.Fatalln(err)
	}

	hyenaLogger.Println("response write history of " + name)
	fmt.Fprintln(os.Stdout, "response write history of "+name)
}
