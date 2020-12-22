#!/bin/bash

export PX4_SIM_MODEL=ssrc_fog_x
source /usr/share/gazebo/setup.sh

# setup Gazebo env and update package path
export GAZEBO_PLUGIN_PATH=$GAZEBO_PLUGIN_PATH:/data/plugins
export GAZEBO_MODEL_PATH=$GAZEBO_MODEL_PATH:/data/models
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/data/plugins

echo "Starting gazebo"
IP_ADDR=$(hostname -I)
GAZEBOIP=${IP_ADDR} GAZEBOMASTER_URI=${IP_ADDR}:11345 gzserver /data/worlds/empty.world --verbose
