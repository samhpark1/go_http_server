package router

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/samhpark1/go_http_server/server"
)

type HandlerFunc func(req *server.Request) *server.Response

type Router struct {
	routes map[string]map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		make(map[string]map[string]HandlerFunc),
	}
}

func ensureDir(path string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s exists but is not a directory", path)
	}
	return nil
}

func (r *Router) AddRoute(method string, route string, handler HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][route] = handler

	//check to see if directory exists if not create
	err := ensureDir(route)
	if err != nil {
		fmt.Println(err)
	}

}

func (r *Router) Server(req *server.Request) *server.Response {
	var reqRoute string

	pattern := `.+\..+$`
	re := regexp.MustCompile(pattern)

	last := req.Path[len(req.Path)-1]

	if re.MatchString(last) {
		reqRoute = strings.Join(req.Path[:len(req.Path)-1], "/")
	} else {
		reqRoute = strings.Join(req.Path, "/")
	}

	handler, exists := r.routes[req.Method][reqRoute]

	if !exists {
		found := false
		for _, methodRoutes := range r.routes {
			if _, ok := methodRoutes[reqRoute]; ok {
				found = true
				break
			}
		}
		if found {
			return HandleNotAllowed(req)
		}
		return HandleNotFound(req)
	}

	for _, val := range req.Path {
		if val == "..." {
			return HandleNotAllowed(req)
		}
	}

	return handler(req)

	//check for not allowed method and route combos and call 405 handler

}

func HandleNotFound(req *server.Request) *server.Response {
	resp := server.CreateResponse(404, req.Version, "Not Found", make(map[string]string), make([]byte, 0))
	return resp
}

func HandleNotAllowed(req *server.Request) *server.Response {
	resp := server.CreateResponse(405, req.Version, "Not Allowed", make(map[string]string), make([]byte, 0))
	return resp
}
