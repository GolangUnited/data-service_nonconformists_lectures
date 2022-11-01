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

func main() {

	db_host := os.Getenv("LECTURES_DB_HOST")
	db_port := os.Getenv("LECTURES_DB_PORT")
	db_user := os.Getenv("LECTURES_DB_USER")
	db_password := os.Getenv("LECTURES_DB_PASSWORD")
	db_name := os.Getenv("LECTURES_DB_NAME")

	db, err := database.New(db_host, db_port, db_user, db_password, db_name)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.AutoMigrate(&models.Lecture{})
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	lecture := &lecture.Lecture{}

	api.RegisterLectureServer(s, lecture)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
