all:
	@echo "Hello"
	
protogen: 	
	protoc -I api/proto --go_out=. --go-grpc_out=. api/proto/lecture.proto

build:
	go build -v -o ./server ./cmd/main.go 

run:
	go run ./cmd/main.go

db-run: db-remove
	docker run --name postgres --network my-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=postgres -p 5432:5432 -d postgres:alpine || true

db-start:
	docker start postgres || true

db-stop:
	docker stop postgres || true

db-remove: db-stop
	docker rm postgres || true

app-start: app-stop
	docker build -t golang-united-lectures .
	docker run --name golang-united-lectures --network my-network -p 8080:8080 -d --rm -e LECTURES_DB_HOST=postgres -e LECTURES_DB_PORT -e LECTURES_DB_USER -e LECTURES_DB_PASSWORD -e LECTURES_DB_NAME golang-united-lectures

app-stop:
	docker stop golang-united-lectures || true
	docker rm golang-united-lectures || true
	docker rmi golang-united-lectures || true	

app-shell:
	docker exec -it golang-united-lectures /bin/sh