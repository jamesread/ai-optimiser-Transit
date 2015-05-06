package main

import (
"fmt"
"libtransit"
)

func main() {
	env := libtransit.Environment{Title: "Default Transit System"}

	fmt.Println("Environment: ", env.Title)
}
