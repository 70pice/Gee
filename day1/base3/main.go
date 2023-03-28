package main

import (
	"fmt"
	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(context *gee.Context) {
		fmt.Fprintln(context.Writer, "this is /")
	})

	r.GET("/hello", func(context *gee.Context) {
		fmt.Fprintln(context.Writer, "this is /hello")
	})
	r.Run(":9999")
}
