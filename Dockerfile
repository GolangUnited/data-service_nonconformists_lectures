FROM golang:alpine

WORKDIR /go/src/golang-united-lectures  

COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o ./server ./cmd/main.go

EXPOSE 8080

CMD ["/go/src/golang-united-lectures/server"]