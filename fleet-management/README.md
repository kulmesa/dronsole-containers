# fleet-management container

## Building and running container

Build and tag container
```
docker build -t tii-fleet-management .
```

Run container in docker
```
docker run --rm -it -p 8082:8082 tii-fleet-management <mqtt-broker-address>
```

## Creating fleet and assigning drones

Create fleet called "Bravo fleet" with id "bravo"
```
curl -d '{"slug":"bravo","name":"Bravo fleet"}' localhost:8082/fleets
```

Assign drone "drone-313" to fleet "bravo"
```
curl -d '{"device_id":"drone-313"}' localhost:8082/fleets/bravo/drones
```

## Listing fleets

```
curl localhost:8082/fleets
```
