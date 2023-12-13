FROM golang:1.18.10-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN rm -rf vendor/* bin/*

RUN go clean -mod=mod
RUN go mod tidy
RUN go mod download && go mod verify

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /library-api

EXPOSE 8080

CMD ["/library-api"]
