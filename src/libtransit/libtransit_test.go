package libtransit

import (
"testing"
)

func TestSanity(t *testing.T) {
	t.Log("My sanity looks okay");
}

func TestSimpleEnvironment(t *testing.T) {
	env := NewEnvironment("Simple environment")
	env.addLift();
	env.addLift();
	env.addLift();

	if env.numLifts() != 3 {
		t.Errorf("Lift count is", env.numLifts())
	}

	env.addFloor();
	env.addFloor();
	env.addFloor();
	env.addFloor();
	env.addFloor();

	if env.numFloors() != 5 {
		t.Errorf("Floor count is", env.numFloors())
	}
}


