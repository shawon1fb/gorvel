package routes

import (
	"net/http"

	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
)

// interface

type RouteInterface interface {
	Register(base *[]Routing) interface{}
	Explain(base *[]Routing) interface{}
}

// structs

type Middlewares []string
type Queries []string
type Method []string
type Resources map[string]Handler
type Handler func(engines.EngineContract) *errors.AppError
type Routing struct {
	Name        string
	Path        string
	Prefix      bool
	Method      []string
	Queries     Queries
	Handler     Handler
	Resources   map[string]Handler
	Middlewares []string
	Static      string
}

// helpers

func resources(route Routing) []Routing {
	routings := []Routing{}
	for action, handler := range route.Resources {
		switch action {
		case "index":
			routings = append(routings, Routing{
				Path:        route.Path,
				Handler:     handler,
				Method:      Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".lists",
			})
		case "create":
			routings = append(routings, Routing{
				Path:        route.Path + "/create",
				Handler:     handler,
				Method:      Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".create",
			})
		case "store":
			routings = append(routings, Routing{
				Path:        route.Path,
				Handler:     handler,
				Method:      Method{"POST"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".store",
			})
		case "show":
			routings = append(routings, Routing{
				Path:        route.Path + "/{id:[0-9]+}",
				Handler:     handler,
				Method:      Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".show",
			})
		case "edit":
			routings = append(routings, Routing{
				Path:        route.Path + "/{id:[0-9]+}/edit",
				Handler:     handler,
				Method:      Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".edit",
			})
		case "update":
			routings = append(
				routings,
				Routing{
					Path:        route.Path + "/{id:[0-9]+}",
					Handler:     handler,
					Method:      Method{"PUT"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".update",
				},
				Routing{
					Path:        route.Path + "/{id:[0-9]+}/update",
					Handler:     handler,
					Method:      Method{"POST"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".update.via-post",
				},
			)
		case "destroy":
			routings = append(
				routings,
				Routing{
					Path:        route.Path + "/{id:[0-9]+}",
					Handler:     handler,
					Method:      Method{"DELETE"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".destroy",
				},
				Routing{
					Path:        route.Path + "/{id:[0-9]+}/delete",
					Handler:     handler,
					Method:      Method{"POST"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".destroy.via-post",
				},
			)
		}
	}
	return routings
}

func getMethods(methods Method) Method {
	if len(methods) == 0 {
		methods = Method{http.MethodGet}
	}
	return methods
}
