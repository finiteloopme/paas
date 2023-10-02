package services

import (
	"context"

	log "github.com/finiteloopme/goutils/pkg/log"

	infrapb "github.com/finiteloopme/paas/api/infra/gen/proto/go/me/finiteloop/v1beta1"
)

type server struct {
	infrapb.InfraResourcesServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) ProcessFolder(ctx context.Context, in *infrapb.ProcessFolderRequest) *infrapb.ProcessFolderResponse {
	log.Info("Request obj: " + in.String())
	return &infrapb.ProcessFolderResponse{}
}

func ProcessFolder()
