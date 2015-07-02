package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
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

// Listen start web app
func Listen(port string) {
	staticDir := path.Join(os.Getenv("GOPATH"), "src/github.com/sosuke-k/hyena/root/static")
	fmt.Fprintln(os.Stdout, staticDir)
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", newRouter())
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":"+port, nil)
}
