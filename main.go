package main

import (
	"github.com/dillonhafer/otd/on_this_day"
)

func main() {
	println(otd.RandomEvent(otd.Events()))
}
