FROM golang:alpine as build
WORKDIR /go/src/golang-united-lectures  
COPY . .
RUN go mod download
RUN go build -o ./server ./cmd/main.go

FROM alpine
WORKDIR .
COPY --from=build /go/src/golang-united-lectures/server ./server
CMD ["./server"]