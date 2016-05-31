package main

import (
	"github.com/kataras/iris"
	"github.com/vishr/web-framework-benchmark/common"
	"fmt"
)

func main() {
	e := iris.New()

	for _, r := range common.DynamicRoutes {
		switch r.Method {
		case "GET":
			e.Get(r.Path, dynamicRoute)
		case "POST":
			e.Post(r.Path, dynamicRoute)
		case "PUT":
			e.Put(r.Path, dynamicRoute)
		case "PATCH":
			e.Patch(r.Path, dynamicRoute)
		case "DELETE":
			e.Delete(r.Path, dynamicRoute)
		default:
			panic("method not supported")
		}
	}

	e.Listen(":8080")
}

func dynamicRoute(c *iris.Context) {
	c.WriteString(fmt.Sprintf("team: %s, user: %s", c.Param("id"), c.Param("user")))
}
