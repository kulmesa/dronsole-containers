# Dronsole containers

This repository contains containers used with `dronsole` command line utility.

## Running containers with docker

All containers can be run locally with docker. This method is normally used during development. The `dronsole` utility will use `minikube` environment to run these.

### Quickstart

Build containers
```
docker build -t tii-mqtt-server mqtt-server/
docker build -t tii-gzserver gzserver/
docker build -t tii-gzweb gzweb/
docker build -t tii-fog-drone fog-drone/
```

Run simulation containers with stdin and tty.
```
docker run --rm -it -p 8883:8883 tii-mqtt-server
docker run --rm -it -v <gazebo-data-location>:/data -p 8081:8081 -p 11345:11345 tii-gzserver
curl -d '' localhost:8081/simulation/start
docker run --rm -it -v <gazebo-data-location>:/data -p 8080:8080 tii-gzweb <host.docker.internal>:11345
open http://localhost:8080/
```

Generate drone identity
```
openssl req -x509 -newkey rsa:2048 -keyout drone_identity_private.pem -nodes -out drone_identity_cert.pem -subj "/CN=unused"
```

Run drone container and add to simulation
```
docker run --rm -it -p 4560:4560 -p 14560:14560/udp -e DRONE_DEVICE_ID="deviceid" -e DRONE_IDENTITY_KEY="$(cat drone_identity_private.pem)" -e MQTT_BROKER_ADDRESS="tcp://<host.docker.internal>:8883" tii-fog-drone

curl -d '{"drone_location":"local","device_id":"deviceid","mavlink_address":"<host.docker.internal>","mavlink_tcp_port":4560,"mavlink_udp_port":14560,"pos_x":0,"pos_y":0}' localhost:8081/simulation/drones
```

Send commands to drone with MQTT
```
mosquitto_pub -h localhost -p 8883 -t "/devices/<deviceid>/commands/control" -m '{"Command":"takeoff"}'
mosquitto_pub -h localhost -p 8883 -t "/devices/<deviceid>/commands/control" -m '{"Command":"land"}'
```

More information can be found from each containers own README file.

