package pm

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/bitly/go-simplejson"
)

func Init(configPath string) {
	hyenaPath := filepath.Dir(configPath)
	if err := os.MkdirAll(hyenaPath, 0777); err != nil {
		log.Fatal(err)
	}
	js := simplejson.New()
	js.Set("length", 0)
	js.Set("projects", simplejson.New())
	w, err := os.Create(configPath)
	if err != nil {
		log.Fatal(err)
		log.Fatal(configPath)
	}
	defer w.Close()
	o, _ := js.EncodePretty()
	w.Write(o)
}

func Load(configPath string) (projects []string) {

	r, err := os.Open(configPath)
	if err != nil {
		log.Fatal(configPath + " is not found")
	} else {
		js, err := simplejson.NewFromReader(r)
		if err != nil {
			log.Fatal("cannot read " + configPath)
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

func Add(configPath string, newProject string) {
	projects := Load(configPath)
	projects = append(projects, newProject)
	js := simplejson.New()
	projectJson := simplejson.New()

	for i, v := range projects {
		projectJson.Set(strconv.Itoa(i), v)
	}
	js.Set("length", len(projects))
	js.Set("projects", projectJson)
	w, err := os.Create(configPath)
	if err != nil {
		log.Fatal(err)
		log.Fatal(configPath)
	}
	defer w.Close()
	o, _ := js.EncodePretty()
	w.Write(o)

	hyenaPath := filepath.Dir(configPath)
	projectDir := path.Join(hyenaPath, newProject)
	if err := os.MkdirAll(projectDir, 0777); err != nil {
		log.Fatal(err)
	}
}
