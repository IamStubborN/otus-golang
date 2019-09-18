package server

import (
	"context"

	"github.com/IamStubborN/otus-golang/2019_08_22_homework_15/calendar/service/event"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
)

type Server struct {
	EService event.EvInterface
}

func NewServer() Server {
	return Server{
		EService: event.NewService(),
	}
}

func (s *Server) Create(ctx context.Context, req *event.CreateRequest) (*event.CreateResponse, error) {
	ev, err := s.EService.Create(req.Event)
	if err != nil {
		return nil, err
	}

	return &event.CreateResponse{
		Event: ev,
	}, nil
}

func (s *Server) Read(ctx context.Context, req *event.ReadRequest) (*event.ReadResponse, error) {
	ev, err := s.EService.Read(req.Event_ID)
	if err != nil {
		return nil, err
	}

	return &event.ReadResponse{Event: ev}, nil
}

func (s *Server) Update(ctx context.Context, req *event.UpdateRequest) (*event.UpdateResponse, error) {
	updated, err := s.EService.Update(req.Event)
	if err != nil {
		return nil, err
	}

	return &event.UpdateResponse{
		Updated: updated,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *event.DeleteRequest) (*event.DeleteResponse, error) {
	deleted, err := s.EService.Delete(req.Event_ID)
	if err != nil {
		return nil, err
	}

	return &event.DeleteResponse{
		Deleted: deleted,
	}, nil
}

func (s *Server) Run(logger *logrus.Logger) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	event.RegisterEventServiceServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve %v", err)
	}
}
