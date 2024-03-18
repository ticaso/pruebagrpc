package handler

import (
	"context"

	pruebagrpc "github.com/ticaso/pruebagrpc/api"
	repository "github.com/ticaso/pruebagrpc/repository"
)

type Server struct {
	pruebagrpc.UnimplementedAspiranteServiceServer
	repo *repository.Repository
}

func NewServer(repo *repository.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) GetAspirante(ctx context.Context, req *pruebagrpc.GetAspiranteRequest) (*pruebagrpc.AspiranteResponse, error) {
	aspirante, err := s.repo.GetAspirante(int64(req.Id))
	if err != nil {
		return nil, err
	}

	return &pruebagrpc.AspiranteResponse{Aspirante: aspirante}, nil
}

func (s *Server) GetAspirantes(ctx context.Context, req *pruebagrpc.AspiranteRequest) (*pruebagrpc.AspirantesResponse, error) {
	aspirantes, err := s.repo.GetAspirantes()
	if err != nil {
		return nil, err
	}

	return &pruebagrpc.AspirantesResponse{Aspirantes: aspirantes}, nil
}
