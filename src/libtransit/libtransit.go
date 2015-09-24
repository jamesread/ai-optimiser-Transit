package libtransit

import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
)

type person struct {
	name string
}

type floor struct {
	people []person
	level int
	environment *Environment
}

type lift struct {
	people []person
	currentFloor int
	movingToFloor int
	movingDirection string
	id int
	environment *Environment
	moving bool
}

func (lift) new() *lift {
	var l = &lift{}
	l.moving = false

	return l;
}

func (l *lift) move() {
	fmt.Println("lift", l.id, "move()")

	if l.currentFloor == l.environment.numFloors() {
		l.movingDirection = "down"
	} else if (l.currentFloor == 1) {
		l.movingDirection = "up";
	}

	l.moving = true

	switch (l.movingDirection) {
		case "up":
			l.currentFloor++;
			l.environment.listener.OnLiftArrived(l);
			break

		default:
			l.currentFloor--;
			l.environment.listener.OnLiftArrived(l);
			break

	}

	l.moving = false

	fmt.Println("lift", l.id, "is now on floor", l.currentFloor)
}

type EnvironmentListener struct {
	environment *Environment
}

func (*EnvironmentListener) OnLiftArrived(lift *lift) {
}

func (*EnvironmentListener) OnPassengerAdded(floor *floor, person *person) {
}

type Environment struct {
	Title string
	floors []floor
	lifts []lift
	listener *EnvironmentListener
}

func NewEnvironment(envTitle string) *Environment {
	e := &Environment{Title: envTitle }

	return e;
}

func (env *Environment) Simulate(count int) {
	for i := 0; i < count; i++ {
		fmt.Println("sim");

		for i, lift := range env.lifts {
			lift = env.lifts[i]
			fmt.Println(i, lift.currentFloor)

			if lift.moving == false {
				go env.lifts[i].move()
			}
		}

		env.RandomFloor().addPassenger(1);

		env.Process();

		fmt.Println(env.toString())
		time.Sleep(1 * time.Second)
	}
}

func (env *Environment) Process() {
	
}

func (env *Environment) toString() string {
	var output string = ""
	
	for i := len(env.floors); i != 0; i-- {
		floor := env.floors[i - 1]

		output += strconv.Itoa(floor.level) + ": - (" + strconv.Itoa(len(floor.people)) + ") ---- "	

		for j := 0; j < len(env.lifts); j++ {
			lift := env.lifts[j]

			if lift.currentFloor == floor.level {
				output += "[" + strconv.Itoa(lift.id) + "] "
			} else {
				output += " |  "
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

func (env *Environment) AddLift() *lift {
	l := &lift{currentFloor: 1, id: env.numLifts() + 1, environment: env}

	env.lifts = append(env.lifts, *l);

	return l;
}

func (env *Environment) AddFloor()  {
	f := floor{level: env.numFloors() + 1}
	f.environment = env

	env.floors = append(env.floors, f);
}

func (env *Environment) RandomFloor() *floor {
	floorIndex := rand.Intn(len(env.floors));
	return &env.floors[floorIndex];	
}

func (f *floor) addPassenger(count int) {
	p := person{name: "untitled"}

	f.people = append(f.people, p)
	f.environment.listener.OnPassengerAdded(f, &p);
}
