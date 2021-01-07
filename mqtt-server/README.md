# mqtt-server

Simple mqtt-server which is used in local simulation environments

## Build

```
docker build -t tii-mqtt-server .
```

## Run

```
docker run --rm -it -p 8883:8883 tii-mqtt-server
```
