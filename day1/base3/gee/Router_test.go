package gee

import (
	"fmt"
	"testing"
)

func helpFunction() *router {
	testRouter := &router{}
	//testRouter.roots =
	testRouter.roots = make(map[string]*node)
	testRouter.handlers = make(map[string]HandlerFunc)
	return testRouter
}
func Test_paresePattern(t *testing.T) {
	var pathOne string = "/one/two/three"
	patternOne := parsePattern(pathOne)

	var pathTwo string = "/one/*/three"
	patternTwo := parsePattern(pathTwo)

	fmt.Println(patternOne)
	fmt.Println(patternTwo)
}

func Test_addRouter(t *testing.T) {
	testRouter := helpFunction()
	//testRouter.roots["GET"] = &node{}
	testRouter.addRoute("GET", "/one/two/three", func(context *Context) {
		fmt.Println("testing")
	})

	testRouter.addRoute("GET", "/one/two/four", func(context *Context) {
		fmt.Println("testing")
	})
}

func Test_getRouter(t *testing.T) {
	testRouter := helpFunction()
	//testRouter.roots["GET"] = &node{}
	testRouter.addRoute("GET", "/one/two/three", func(context *Context) {
		fmt.Println("testing")
	})

	testRouter.addRoute("GET", "/one/two/four", func(context *Context) {
		fmt.Println("testing")
	})

	testRouter.getRoute("GET", "/one/two/three")
}