FROM golang:1.19

WORKDIR /usr/src/app

COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app
EXPOSE 5000
CMD ["app"]

