# select image
FROM golang:1.14.1-alpine
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go mod vendor
RUN go build -o main ./api
# Command to run the executable
CMD ["/api"]
