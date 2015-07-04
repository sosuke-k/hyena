package pm

import (
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/sosuke-k/hyena/util/log"
)

// Project struct
type Project struct {
	Name string `json:"name"`
}

// Projects is Project slice
type Projects []Project

// Init creates a config file and if a parent directory not exists, also creates it.
func Init(configPath string) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to create config file " + configPath)
	hyenaLogger.Println("even if already exists, overwride config file")
	hyenaPath := filepath.Dir(configPath)
	if err := os.MkdirAll(hyenaPath, 0777); err != nil {
		hyenaLogger.Fatalln(err)
	}
	js := simplejson.New()
	js.Set("length", 0)
	js.Set("projects", simplejson.New())
	w, err := os.Create(configPath)
	if err != nil {
		hyenaLogger.Fatalln(err)
	}
	defer w.Close()
	o, _ := js.EncodePretty()
	w.Write(o)
}

// Load returns the string array of projects' name by the path of config file
func Load(configPath string) (projects []string) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to load projects list")

	r, err := os.Open(configPath)
	if err != nil {
		hyenaLogger.Println(configPath + " is not found")
	} else {
		js, err := simplejson.NewFromReader(r)
		if err != nil {
			hyenaLogger.Println("cannot read " + configPath)
		} else {
			length, _ := js.Get("length").Int()
			for i := 0; i < length; i++ {
				project, _ := js.Get("projects").Get(strconv.Itoa(i)).String()
				projects = append(projects, project)
			}
		}
	}
	return
}

// Add adds a project to project list by the path of config file
func Add(configPath string, newProject string) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to add " + newProject + " project")
	projects := Load(configPath)
	projects = append(projects, newProject)
	js := simplejson.New()
	projectJSON := simplejson.New()

	for i, v := range projects {
		projectJSON.Set(strconv.Itoa(i), v)
	}
	js.Set("length", len(projects))
	js.Set("projects", projectJSON)
	w, err := os.Create(configPath)
	if err != nil {
		hyenaLogger.Fatalln(err)
	}
	defer w.Close()
	o, _ := js.EncodePretty()
	w.Write(o)

	hyenaPath := filepath.Dir(configPath)
	projectDir := path.Join(hyenaPath, newProject)
	if err := os.MkdirAll(projectDir, 0777); err != nil {
		hyenaLogger.Fatalln(err)
	}
}

// Delete deletes a project from project list by the path of config file
func Delete(configPath string, targetProject string) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to delete " + targetProject + " project")

	projects := Load(configPath)
	for i, name := range projects {
		if name == targetProject {
			projects = append(projects[:i], projects[i+1:]...)
			break
		}
	}

	js := simplejson.New()
	projectJSON := simplejson.New()
	for i, v := range projects {
		projectJSON.Set(strconv.Itoa(i), v)
	}
	js.Set("length", len(projects))
	js.Set("projects", projectJSON)
	w, err := os.Create(configPath)
	if err != nil {
		hyenaLogger.Fatalln(err)
	}
	defer w.Close()
	o, _ := js.EncodePretty()
	w.Write(o)
}
