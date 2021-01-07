# gzserver container

This container will provide single simulation environment running Gazebo Server (gzserver) with a service providing API to control the simulation.

## Building and running container

Build and tag container
```
docker build -t tii-gzserver .
```

To run the simulation you must have the gazebo data available so it can be mounted to the server.

Run container in docker
```
docker run --rm -it -v <gazebo-data-location>:/data -p 8081:8081 -p 11345:11345 tii-gzserver
```

## Starting and stopping the simulation

The Gazebo simulation can be started and stopped by calling the service running in port 8081
```
curl -d '' localhost:8081/simulation/start
curl -d '' localhost:8081/simulation/stop
```

## Adding drone to the simulation

Before adding the drone to the simulation you should have the px4 and other software running.

Add the drone to the simulation
```
curl -d '{"drone_location":"local","device_id":"deviceid","mavlink_address":"host.docker.internal","mavlink_tcp_port":4560,"mavlink_udp_port":14560,"pos_x":0,"pos_y":0}' localhost:8081/simulation/drones
```

## List drones in simulation

```
curl localhost:8081/simulation/drones
```

## Building and running locally

```
go build -o ./gzserver-api .
./gzserver-api
```
