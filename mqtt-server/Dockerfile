FROM golang:latest as builder
RUN mkdir /mqtt-server
WORKDIR /mqtt-server/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go ./
RUN GOOS=linux GOARCH=amd64 go build -o mqtt-server .


FROM golang:latest
COPY --from=builder /mqtt-server/mqtt-server /bin/mqtt-server
EXPOSE 8883
ENTRYPOINT ["/bin/mqtt-server"]
