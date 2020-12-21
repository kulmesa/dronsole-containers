# Fog drone container image

## Build the fog drone image

```
docker build -t tii-fog-drone .
```

## Running the fog drone locally

```
export DRONE_DEVICE_ID="device_id"
export DRONE_IDENTITY_KEY="$(cat rsa_private.pem)"
export MQTT_BROKER_ADDRESS="tcp://ip:port"
# export MQTT_BROKER_ADDRESS="ssl://mqtt.googleapis.com:8883"

# run container and forward environment variables
docker run -it --network=host \
    -e DRONE_DEVICE_ID \
    -e DRONE_IDENTITY_KEY \
    -e MQTT_BROKER_ADDRESS \
    tii-fog-drone
```
