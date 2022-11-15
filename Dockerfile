FROM golang:latest
RUN apk update
RUN apk install git
RUN go get github.com/Unleash/unleash-client-go
COPY catweb.go /go

RUN go build catweb.go

FROM alpine:latest
COPY --from=0 /go/catweb /
COPY templates/index.html /
COPY static/* /static/
EXPOSE 5000
CMD /catweb
