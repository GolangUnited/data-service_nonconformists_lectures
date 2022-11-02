package main

import (
	"golang-united-lectures/internal/api"
	"golang-united-lectures/internal/database"
	"golang-united-lectures/internal/lecture"
	"golang-united-lectures/internal/models"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

const protocol_tcp = "tcp"
const port_8080 = ":8080"

func main() {

	dbHost := os.Getenv("LECTURES_DB_HOST")
	dbPort := os.Getenv("LECTURES_DB_PORT")
	dbUser := os.Getenv("LECTURES_DB_USER")
	dbPassword := os.Getenv("LECTURES_DB_PASSWORD")
	dbName := os.Getenv("LECTURES_DB_NAME")

	err := database.New(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	err = database.DB.AutoMigrate(&models.Lecture{})
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	lectureServer := &lecture.Lecture{}

	api.RegisterLectureServer(grpcServer, lectureServer)

	listener, err := net.Listen(protocol_tcp, port_8080)
	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
