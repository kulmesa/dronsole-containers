FROM golang:latest as builder
RUN mkdir /simulator
WORKDIR /simulator/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go ./
RUN GOOS=linux GOARCH=amd64 go build -o /mqtt-server .


FROM golang:latest
COPY --from=builder /mqtt-server /mqtt-server
EXPOSE 8883
ENTRYPOINT ["/mqtt-server"]
