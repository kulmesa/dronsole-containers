#!/bin/bash

source /opt/ros/foxy/setup.bash
echo "Start Mavlink Router"
mavlink-routerd >/fog-drone/mav_router_out.log 2>/fog-drone/mav_router_err.log &
echo "Start Mavlink control"
ros2 launch px4_mavlink_ctrl mavlink_ctrl.launch >/fog-drone/mav_ctrl_out.log 2>/fog-drone/mav_ctrl_err.log &
echo "Start Micrortps_agent"
micrortps_agent -t UDP >/fog-drone/urtps_out.log 2>/fog-drone/urtps_err.log &
echo "Start Communication link"
communication_link -device_id "$DRONE_DEVICE_ID" -mqtt_broker "$MQTT_BROKER_ADDRESS"
