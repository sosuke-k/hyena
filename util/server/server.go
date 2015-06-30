package server

import (
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
	http.ListenAndServe(":"+port, newRouter())
}
