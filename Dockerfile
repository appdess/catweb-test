FROM golang:alpine3.16
COPY catweb.go /go
COPY go.mod /go
COPY go.sum /go
RUN go get github.com/Unleash/unleash-client-go
RUN go build catweb.go

FROM alpine:latest
COPY --from=0 /go/catweb /
COPY templates/index.html /
COPY static/* /static/
EXPOSE 5000
CMD /catweb
