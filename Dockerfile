FROM golang:1.13
WORKDIR /go/src
COPY go.mod .
COPY go.sum .
RUN go mod download
EXPOSE 9000
CMD ["go", "run", "/go/src/main.go"]