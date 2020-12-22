# gzserver container

This container will provide single simulation environment running Gazebo Server (gzserver) with a service providing API to control the simulation.

## Building and running container

Build and tag container
```
docker build -t tii-gzserver .
```

Run it locally
```
docker run -p 8081:8081 tii-gzserver
```


## Building and running locally

```
go build -o ./gzserver-api .
./gzserver-api
```
