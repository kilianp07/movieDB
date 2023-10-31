# Build Golang binary
FROM golang:1.20.2 AS build-golang

WORKDIR /Users/kilianpeyron/Repos/movieDB

COPY . .
RUN go get -v && go build -v -o /usr/local/bin/movieDB

EXPOSE 8080
CMD ["movieDB"]
