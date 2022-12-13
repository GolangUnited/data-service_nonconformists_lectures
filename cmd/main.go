package main

import (
	"fmt"
	"golang-united-lectures/config"
	"golang-united-lectures/pkg/api"
	"golang-united-lectures/pkg/database"
	"golang-united-lectures/pkg/repositories"
	"golang-united-lectures/pkg/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	err = database.DB.AutoMigrate(&repositories.Lecture{})
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	lectureServer := &service.Lecture{}

	api.RegisterLectureServer(grpcServer, lectureServer)

	listener, err := net.Listen(config.PROTOCOL_TCP, fmt.Sprintf(":%s", config.PORT_8080))
	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
