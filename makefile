all:
	@echo "You can use protogen, run, start, stop or shell parameters"
	
protogen: 	
	protoc -I api/proto --go_out=. --go-grpc_out=. api/proto/lecture.proto

run:
	go run ./cmd/main.go

start: stop
	docker build -t golang-united-lectures .
	docker run --name golang-united-lectures --network my-network -p 8080:8080 -d -e LECTURES_DB_HOST=postgres -e LECTURES_DB_PORT -e LECTURES_DB_USER -e LECTURES_DB_PASSWORD -e LECTURES_DB_DATABASE golang-united-lectures

stop:
	docker stop golang-united-lectures || true
	docker rm golang-united-lectures || true
	docker rmi golang-united-lectures || true	

shell:
	@docker exec -it golang-united-lectures /bin/sh