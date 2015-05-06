package libtransit

import (
"testing"
)

func TestSanity(t *testing.T) {
	t.Log("My sanity looks okay");
}

func TestSimpleEnvironment(t *testing.T) {
	env := NewEnvironment("Simple environment")
	env.AddLift();
	env.AddLift();
	env.AddLift();

	if env.numLifts() != 3 {
		t.Errorf("Lift count is", env.numLifts())
	}

	env.AddFloor();
	env.AddFloor();
	env.AddFloor();
	env.AddFloor();
	env.AddFloor();

	if env.numFloors() != 5 {
		t.Errorf("Floor count is", env.numFloors())
	}
}


