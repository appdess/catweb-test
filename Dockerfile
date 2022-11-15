FROM golang:1.19
WORKDIR /go
COPY catweb.go /go
RUN go mod download && go mod verify
RUN go build catweb.go

FROM alpine:latest
COPY --from=0 /go/catweb /
COPY templates/index.html /
COPY static/* /static/
EXPOSE 5000
CMD /catweb
