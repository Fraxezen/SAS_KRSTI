package main

import (
	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/pkg/msgs/control_msgs"
	"os"
	"os/signal"
	"rosaya/pkg/log"
	"time"
)

func main() {
	args := os.Args

	if len(args) > 1 {
		nodeConf := goroslib.NodeConf{
			Namespace:     "/" + args[1],
			Name:          "controller_actionserver",
			MasterAddress: "localhost:11311",
		}

	RETRY:
		n, err := goroslib.NewNode(nodeConf)

		if err != nil {
			log.Errorf("error: %v", err)
			time.Sleep(time.Second)
			goto RETRY
		}
		log.Infof("node %s/%s is up", nodeConf.Namespace, nodeConf.Name)
		defer n.Close()

		sasConf := goroslib.SimpleActionServerConf{
			Node:   n,
			Name:   "follow_joint_trajectory",
			Action: &control_msgs.FollowJointTrajectoryAction{},
		}
		onExecute := func(sas *goroslib.SimpleActionServer, goal *control_msgs.FollowJointTrajectoryActionGoal) {
			log.Infof("onExecute: goal: %v", len(goal.Trajectory.Points))
		}
		sasConf.OnExecute = onExecute

		// create a simple action server
		sas, err := goroslib.NewSimpleActionServer(sasConf)
		if err != nil {
			panic(err)
		}
		defer sas.Close()

		// wait for CTRL-C
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
	} else {
		log.Errorf("need an action name as argument")
	}
}
