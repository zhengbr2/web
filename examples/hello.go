package main

import (
	"hoise_web"
)

func hello(val string) string { return "hello " + val + "\n" }
func foo(val string) string   { return "bar \n" }

func main() {

	web.Get("/(.*)", hello)
	//web.Get("/foo", foo)
	web.Run("0.0.0.0:9999")
}
