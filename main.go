package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	pruebagrpc "github.com/ticaso/pruebagrpc/api"
	"google.golang.org/grpc"
)

// Asegúrate de que la implementación del servidor utiliza directamente los nombres
// de los servicios y mensajes definidos en los archivos .pb.go sin un prefijo.

func connectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/aspirante")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	} else {
		log.Println("Connected to the database.")
	}

	return db
}

type server struct {
	pruebagrpc.UnimplementedAspiranteServiceServer
	db *sql.DB
}

func (s *server) GetAspirante(ctx context.Context, req *pruebagrpc.GetAspiranteRequest) (*pruebagrpc.AspiranteResponse, error) {
	var aspirante pruebagrpc.Aspirante
	err := s.db.QueryRow("SELECT id, nombre, apellido FROM aspirante WHERE id = ?", req.Id).Scan(&aspirante.Id, &aspirante.Nombre, &aspirante.Apellido)
	if err != nil {
		return nil, err
	}

	return &pruebagrpc.AspiranteResponse{Aspirante: &aspirante}, nil
}

func (s *server) GetAspirantes(ctx context.Context, req *pruebagrpc.AspiranteRequest) (*pruebagrpc.AspirantesResponse, error) {
	rows, err := s.db.Query("SELECT id, nombre, apellido FROM aspirante")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aspirantes []*pruebagrpc.Aspirante
	for rows.Next() {
		var a pruebagrpc.Aspirante
		if err := rows.Scan(&a.Id, &a.Nombre, &a.Apellido); err != nil {
			return nil, err
		}
		aspirantes = append(aspirantes, &a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pruebagrpc.AspirantesResponse{Aspirantes: aspirantes}, nil
}

func main() {
	db := connectToDB() // Usa tu función existente para conectar a la base de datos.
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pruebagrpc.RegisterAspiranteServiceServer(s, &server{db: db})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
