# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /app
# Copy Swagger UI files
COPY ./docs /app/docs

COPY go.mod go.sum .
RUN go mod download

COPY . .

# RUN go build  -o /go-docker-demo

EXPOSE 3000

CMD [ "go","run","main.go" ]
