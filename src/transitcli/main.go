package main

import (
	"fmt"
	"libtransit"
	goopt "github.com/droundy/goopt"
)

func main() {
	var simulationSteps = goopt.Int([]string{"-s", "--simulationSteps"}, 10, "Simulation steps.")
	var floorCount = goopt.Int([]string{"-f", "--floorCount"}, 5, "floor count")
	var liftCount = goopt.Int([]string{"-l", "--liftCount"}, 2, "lift count")
	var title = goopt.String([]string{"-t", "--title"}, "title", "title of the environment")

	goopt.Parse(nil)
		
	env := libtransit.Environment{Title: *title}

	for i := 0; i < *floorCount; i++ {
		env.AddFloor();
	}

	for i := 0; i < *liftCount; i++ {
		env.AddLift();
	}

	fmt.Println("Environment: ", env.Title)
	fmt.Printf("%+v\n", env)

	env.Simulate(*simulationSteps);
}
