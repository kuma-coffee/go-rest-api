package main

import (
	"github.com/kuma-coffee/go-rest-api/routes"
)

func main() {
	e := routes.Route()

	e.Logger.Fatal(e.Start(":1323"))
}
