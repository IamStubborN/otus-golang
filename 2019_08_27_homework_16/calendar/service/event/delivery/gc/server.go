package gc

import (
	"context"
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/service/event/repository"

	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/service/event/domain"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
)

type Server struct {
	Storage repository.EvInterface
}

func NewServer(storage repository.EvInterface) Server {
	return Server{
		Storage: storage,
	}
}

func (s *Server) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	ev := &domain.Event{
		ID:          req.Event.ID,
		Name:        req.Event.Name,
		Description: req.Event.Description,
		Date:        req.Event.Date,
	}

	ev, err := s.Storage.Create(ev)
	if err != nil {
		return nil, err
	}

	gEvent := &Event{
		ID:          ev.ID,
		Name:        ev.Name,
		Description: ev.Description,
		Date:        ev.Date,
	}

	return &CreateResponse{
		Event: gEvent,
	}, nil
}

func (s *Server) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	ev, err := s.Storage.Read(req.Event_ID)
	if err != nil {
		return nil, err
	}

	gEvent := &Event{
		ID:          ev.ID,
		Name:        ev.Name,
		Description: ev.Description,
		Date:        ev.Date,
	}

	return &ReadResponse{Event: gEvent}, nil
}

func (s *Server) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	ev := &domain.Event{
		ID:          req.Event.ID,
		Name:        req.Event.Name,
		Description: req.Event.Description,
		Date:        req.Event.Date,
	}

	updated, err := s.Storage.Update(ev)
	if err != nil {
		return nil, err
	}

	return &UpdateResponse{
		Updated: updated,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	deleted, err := s.Storage.Delete(req.Event_ID)
	if err != nil {
		return nil, err
	}

	return &DeleteResponse{
		Deleted: deleted,
	}, nil
}

func (s *Server) Run(logger *logrus.Logger) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Fatalf("failed to listen %v", err)
	}

	gServer := grpc.NewServer()
	reflection.Register(gServer)

	RegisterEventServiceServer(gServer, s)
	if err := gServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve %v", err)
	}
}
