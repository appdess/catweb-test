FROM golang:1.8-alpine
RUN apk update
RUN apk add git
RUN go get github.com/Unleash/unleash-client-go/v3
COPY catweb.go /go
RUN go build catweb.go

FROM alpine:latest
COPY --from=0 /go/catweb /
COPY templates/index.html /
COPY static/* /static/
EXPOSE 5000
CMD /catweb
