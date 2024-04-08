# syntax=docker/dockerfile:1

FROM golang:1.20.3


# Set destination for COPY
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o myapp .

CMD ["./myapp"]



