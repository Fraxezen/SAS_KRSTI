# SAS_KRSTI

## Dependencies
- ros noetic [http://wiki.ros.org/noetic/Installation/Ubuntu](http://wiki.ros.org/noetic/Installation/Ubuntu), or [using docker](#using-docker)
- moveit [install here](https://moveit.ros.org/install/)
- track-ik kinematics solver [install here](https://ros-planning.github.io/moveit_tutorials/doc/trac_ik/trac_ik_tutorial.html)
- [ros-noetic-joint-state-publisher](https://wiki.ros.org/joint_state_publisher)

## Using docker
first install docker [here](https://docs.docker.com/engine/install/ubuntu/) if you don't have it already.

then run the following commands to build the docker image and run the container:
```bash
docker-compose up -d
```
this will build the image and run the container. you can then install the dependencies inside the container by running:
```bash
docker-compose exec -it pocketros /usr/bin/sudo apt -y install ros-noetic-moveit ros-noetic-trac-ik-kinematics-plugin ros-noetic-joint-state-publisher
```
build the package inside the container by running:
```bash
docker-compose exec -it pocketros /bin/bash -c "cd /home/pocketros/catkin_ws && catkin build"
```
and finally run the package by running:
```bash
docker-compose exec -it pocketros /bin/bash -c "source /home/pocketros/catkin_ws/devel/setup.bash && DISPLAY=:0 roslaunch rosaya rosaya.launch"
```
you can now open rviz and see the robot model inside novnc by going to [http://localhost:8080](http://localhost:8080).