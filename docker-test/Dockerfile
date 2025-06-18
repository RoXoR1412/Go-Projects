FROM golang:1.24.4-alpine3.22
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY main.go .
RUN go build -o /godocker
EXPOSE 8080
CMD ["/godocker"]
