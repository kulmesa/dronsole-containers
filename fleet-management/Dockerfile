FROM golang:latest as builder
RUN mkdir /fleet-management
WORKDIR /fleet-management/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go ./
RUN GOOS=linux GOARCH=amd64 go build -o fleet-management .


FROM golang:latest
COPY --from=builder /fleet-management/fleet-management /bin/fleet-management
EXPOSE 8082
ENTRYPOINT ["/bin/fleet-management"]
