package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is Route slice
type Routes []Route

func newRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Home",
		"GET",
		"/",
		homeHandler,
	},
	Route{
		"Projects",
		"GET",
		"/project/{name}",
		projectHandler,
	},
	Route{
		"ProjectsAPI",
		"GET",
		"/api/projects",
		projectListAPIHandler,
	},
	Route{
		"ProjectDiffAPI",
		"GET",
		"/api/project/{name}/diff/{oldCommit}/{newCommit}",
		projectDiffAPIHandler,
	},
}
