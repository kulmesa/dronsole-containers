FROM golang:latest as builder

RUN mkdir /gzserver-api
WORKDIR /gzserver-api/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o gzserver-api .

FROM gazebo:gzserver11-focal

RUN apt-get update -y && apt-get install -y --no-install-recommends \
    gazebo11-plugin-base \
    python3-jinja2 \
    python3-numpy \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir /gzserver-api
WORKDIR /gzserver-api
COPY scripts/launch-gzserver.sh                     scripts/launch-gzserver.sh
COPY scripts/spawn-drone.sh                         scripts/spawn-drone.sh

COPY --from=builder /gzserver-api/gzserver-api /bin/gzserver-api

EXPOSE 8081
EXPOSE 11345

ENTRYPOINT ["/bin/gzserver-api"]
