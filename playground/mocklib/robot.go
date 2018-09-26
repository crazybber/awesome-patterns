package mocklib

import "fmt"

// Robot
type Robot interface {
	SayHi()
}

// ServiceRobot is kind of Robot can offer services
type ServiceRobot struct {
}

func (robot *ServiceRobot) SayHi() {
	fmt.Println("Hi, I'm service robot")
}

// IndustrialRobot is kind of Robot can do some jobs
type IndustrialRobot struct {
}

func (robot *IndustrialRobot) SayHi() {
	fmt.Println("Hi, I'm industrial robot")
}

func StartRobots() {
	robots := initializeRobots()
	makeRobotsSayHi(robots)
}

// initialize all robots
func initializeRobots() []Robot {
	robots := []Robot{
		&ServiceRobot{},
		&IndustrialRobot{},
	}
	return robots
}

// makeRobotsSayHi is used for making robots say hi
func makeRobotsSayHi(robots []Robot) {
	for _, robot := range robots {
		robot.SayHi()
	}
}
