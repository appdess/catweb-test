FROM golang:1.8-alpine
WORKDIR /usr/src/app
COPY go.mod go.sum ./
COPY . .
RUN go build -v -o /usr/local/bin/app ./...
COPY catweb.go /go
RUN go build catweb.go

FROM alpine:latest
COPY --from=0 /go/catweb /
COPY templates/index.html /
COPY static/* /static/
EXPOSE 5000
CMD /catweb
