package server

import (
	"context"
	"log"
	"runtime"
	"shortlink/api"
	"shortlink/shortcut"

	"google.golang.org/grpc/status"
)

// Server implements API gRPC service
type Server struct {
	api.UnimplementedAPIServiceServer
	ShortcutClient shortcut.ShortcutServiceClient
	// More fields for other service clients
}

func gRPCError(err error) error {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[ERROR]: %s:%d %v", fn, line, err)
	errStatus, _ := status.FromError(err)
	return status.Errorf(errStatus.Code(), errStatus.Message())
}

// CreateShortcut handles CreateShortcut requests, passing them to ShortcutService
func (s *Server) CreateShortcut(ctx context.Context, request *api.CreateShortcutRequest) (*api.CreateShortcutResponse, error) {
	internalCreate := &shortcut.InternalCreateShortcutRequest{
		TargetUrl: request.TargetUrl,
	}

	resp, err := s.ShortcutClient.CreateShortcut(context.Background(), internalCreate)
	if err != nil {
		return nil, gRPCError(err)
	}
	return resp, nil
}

// DeleteShortcut handles DeleteShortcut requests, passing them to ShortcutService
func (s *Server) DeleteShortcut(ctx context.Context, request *api.DeleteShortcutRequest) (*api.DeleteShortcutResponse, error) {

	internalDelete := &shortcut.InternalDeleteShortcutRequest{
		UrlToken: request.UrlToken,
	}

	resp, err := s.ShortcutClient.DeleteShortcut(context.Background(), internalDelete)
	if err != nil {
		return nil, gRPCError(err)
	}
	return resp, nil
}

// GetShortcut handles GetShortcut requests, passing them to ShortcutService
func (s *Server) GetShortcut(ctx context.Context, request *api.GetShortcutRequest) (*api.GetShortcutResponse, error) {
	resp, err := s.ShortcutClient.GetShortcut(context.Background(), request)
	if err != nil {
		return nil, gRPCError(err)
	}
	return resp, nil
}
