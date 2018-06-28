package main

import (
	"github.com/timcurless/wings/flightservice/service"
)

func main() {
	flightservice.StartService("8080")
}
