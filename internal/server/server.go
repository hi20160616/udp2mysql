package server

import (
	"context"
)

func (gs *GRPC) Start(ctx context.Context) error {
	return gs.s.Serve(*gs.l)
}

func (gs *GRPC) Stop(ctx context.Context) error {
	gs.s.GracefulStop()
	return nil
}
