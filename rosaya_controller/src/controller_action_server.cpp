#include <ros/ros.h>
#include <actionlib/server/simple_action_server.h>
#include <control_msgs/FollowJointTrajectoryAction.h>
#include <trajectory_msgs/JointTrajectoryPoint.h>

class ControllerActionServer
{
protected:
    ros::NodeHandle nh;
    actionlib::SimpleActionServer<control_msgs::FollowJointTrajectoryAction> as;
    std::string action_name;
    control_msgs::FollowJointTrajectoryActionFeedback feedback;
    control_msgs::FollowJointTrajectoryActionResult result;

public:
    ControllerActionServer(std::string name) : as(nh, name, boost::bind(&ControllerActionServer::executeCB, this, _1), false), action_name(name)
    {
        as.start();
    }

    ~ControllerActionServer(void)
    {
    }

    void executeCB(control_msgs::FollowJointTrajectoryGoalConstPtr const &goal)
    {
        std::stringstream goalTrajectoryPointPositionsString;
        std::vector<trajectory_msgs::JointTrajectoryPoint> goalTrajectoryPoints = goal->trajectory.points;
        std::vector<trajectory_msgs::JointTrajectoryPoint>::iterator goalTrajectoryPointsIterator;
        int i = 0;
        for (goalTrajectoryPointsIterator = goalTrajectoryPoints.begin(); goalTrajectoryPointsIterator < goalTrajectoryPoints.end(); goalTrajectoryPointsIterator++)
        {
            std::stringstream goalTrajectoryPointPositionsStream;
            std::copy(goalTrajectoryPointsIterator->positions.begin(), goalTrajectoryPointsIterator->positions.end(), std::ostream_iterator<double>(goalTrajectoryPointPositionsStream, " "));
            ROS_INFO_STREAM(action_name.c_str() << ": Received New Planning, With Goal Points " << i << ": " << goalTrajectoryPointPositionsStream.str());
            i++;
        }
    }
};

int main(int argc, char **argv)
{
    ros::init(argc, argv, "controller_action_server");

    ControllerActionServer controllerActionServer("follow_joint_trajectory");
    ros::spin();

    return 0;
}