package main

import (
"fmt"
"libtransit"
"flag"
)

func main() {
	var title = flag.String("title", "Untitled Transit System", "The title for the transit system")
	var floorCount = flag.Int("floorCount", 5, "The number of floors")
	var liftCount = flag.Int("liftcount", 2, "The number of lifts")
	var simulate = flag.Bool("Simulate", false, "Start the simulation?")

	flag.Parse()

	env := libtransit.Environment{Title: *title}

	for i := 0; i < *floorCount; i++ {
		env.AddFloor();
	}

	for i := 0; i < *liftCount; i++ {
		env.AddLift();
	}

	fmt.Println("Environment: ", env.Title)
	fmt.Println(env)

	if *simulate {
		env.Simulate();
	}
}
