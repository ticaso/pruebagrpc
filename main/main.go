package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	pruebagrpc "github.com/ticaso/pruebagrpc/api"
	"github.com/ticaso/pruebagrpc/handler"
	"github.com/ticaso/pruebagrpc/repository"

	"google.golang.org/grpc"
)

func connectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/aspirante")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	} else {
		log.Println("Connected to the database.")
	}

	return db
}

func main() {
	db := connectToDB()
	defer db.Close()

	repo := repository.NewRepository(db)
	h := handler.NewServer(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pruebagrpc.RegisterAspiranteServiceServer(s, h)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
