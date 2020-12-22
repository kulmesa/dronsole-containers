# gzweb

## Building the container image

```
docker build -t tii-gzweb .
```

## Running locally

You need to have the models available at local machine

```
export GAZEBO_MODELS="/some/local/directory"
export GAZEBO_SERVER_ADDRESS="some.net.ip.add:port"
docker run -it -p 8080:8080 -v $GAZEBO_MODELS:/data tii-gzweb $GAZEBO_SERVER_ADDRESS
```
