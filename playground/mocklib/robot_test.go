package mocklib

import (
	"github.com/weichou1229/go-patterns/playground/mocklib/mocks"
	"testing"
)

func TestStartRobots(t *testing.T) {
	StartRobots()
}

func TestMakeRobotsSayHi(t *testing.T) {
	// create an instance of our test object
	mockRobotA := new(mocks.Robot)
	mockRobotB := new(mocks.Robot)

	// setup expectations
	mockRobotA.On("SayHi").Return(nil, nil)
	mockRobotB.On("SayHi").Return(nil, nil)

	robots := []Robot{
		mockRobotA,
		mockRobotB,
	}

	// Act
	makeRobotsSayHi(robots)

	// Assert that the expectations were met
	mockRobotA.AssertExpectations(t)
	mockRobotB.AssertExpectations(t)
}
