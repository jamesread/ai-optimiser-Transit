package libtransit

import (
	"fmt"
	"time"
	"strconv"
)

type person struct {
	name string
}

type floor struct {
	people []person
	level int
}

type lift struct {
	people []person
	currentFloor int
	movingToFloor int
	id string

	moving bool
}

func (lift) new() lift {
	var l = lift{}
	l.moving = false

	return l;
}

func (l lift) move() {
	fmt.Println("move()")

	l.moving = true

	time.Sleep(3 * time.Second)
	l.currentFloor++;

	fmt.Println("lift is now on floor:", l.currentFloor)

	l.moving = false
}

type Environment struct {
	Title string
	floors []floor
	lifts []lift
}

func NewEnvironment(envTitle string) *Environment {
	e := &Environment{Title: envTitle }

	return e;
}

func (env *Environment) Simulate() {
	for {
		fmt.Println("sim");

		for _, lift := range env.lifts {
			if lift.moving == false {
				go lift.move()
			}
		}

		fmt.Println(env.toString())
		time.Sleep(1 * time.Second)
	}
}

func (env *Environment) toString() string {
	var output string = ""
	
	for _, floor := range env.floors {
		output += strconv.Itoa(floor.level) + ": - (" + strconv.Itoa(len(floor.people)) + ") ---- "	

		for _, lift := range env.lifts {
			if lift.currentFloor == floor.level {
				output += "#"
			} else {
				output += "_"
			}
		}

		output += "\n"
	}

	return output;
}

func (env *Environment) numLifts() int {
	return len(env.lifts)
}

func (env *Environment) numFloors() int { 
	return len(env.floors)
}

func (env *Environment) AddLift() {
	l := lift{currentFloor: 1}

	env.lifts = append(env.lifts, l);
}

func (env *Environment) AddFloor()  {
	f := floor{level: env.numFloors() + 1}

	env.floors = append(env.floors, f);
}

