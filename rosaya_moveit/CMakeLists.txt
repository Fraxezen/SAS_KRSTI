cmake_minimum_required(VERSION 3.1.3)
project(rosaya_moveit)

find_package(catkin REQUIRED COMPONENTS
  roscpp
  rospy
  std_msgs
  sensor_msgs
  geometry_msgs
  actionlib
  actionlib_msgs
  moveit_msgs
  moveit_core
  moveit_ros_planning
  moveit_ros_planning_interface
  dynamixel_workbench_msgs
  dynamixel_workbench_toolbox
  message_generation
)
find_package(Eigen3 REQUIRED)
find_package(roslaunch)

catkin_package(
  INCLUDE_DIRS include
  CATKIN_DEPENDS roscpp std_msgs sensor_msgs geometry_msgs moveit_msgs  moveit_core moveit_ros_planning moveit_ros_planning_interface message_runtime
  DEPENDS EIGEN3
)

include_directories(
  include
  ${catkin_INCLUDE_DIRS}
  ${EIGEN3_INCLUDE_DIRS}
)



add_executable(trajectory_server src/trajectory_server.cpp)
add_dependencies(trajectory_server ${${PROJECT_NAME}_EXPORTED_TARGETS} ${catkin_EXPORTED_TARGETS})
target_link_libraries(trajectory_server ${catkin_LIBRARIES} ${Eigen3_LIBRARIES})

install(DIRECTORY launch DESTINATION ${CATKIN_PACKAGE_SHARE_DESTINATION}
  PATTERN "setup_assistant.launch" EXCLUDE)
install(DIRECTORY config DESTINATION ${CATKIN_PACKAGE_SHARE_DESTINATION})
install(TARGETS trajectory_server 
  RUNTIME DESTINATION ${CATKIN_PACKAGE_SHARE_DESTINATION}
)
install(DIRECTORY include/${PROJECT_NAME}/
  DESTINATION ${CATKIN_PACKAGE_INCLUDE_DESTINATION}
)
