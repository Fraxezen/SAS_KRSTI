cmake_minimum_required(VERSION 3.0.2)
project(rosaya_controller)

find_package(catkin REQUIRED COMPONENTS roscpp actionlib control_msgs)
find_package(Boost REQUIRED COMPONENTS system)

catkin_package(
	CATKIN_DEPENDS control_msgs
)

include_directories(include ${catkin_INCLUDE_DIRS} ${Boost_INCLUDE_DIRS})

add_executable(controller_action_server src/controller_action_server.cpp)

target_link_libraries(
  controller_action_server
  ${catkin_LIBRARIES}
)