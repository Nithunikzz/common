package outh

import (
	cerbosclient "github.com/Nithunikzz/common/cerbose"
	"github.com/Nithunikzz/common/config"
	oryclient "github.com/Nithunikzz/common/ory"
	"github.com/Nithunikzz/common/ports"
	"github.com/intelops/go-common/logging"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Adapter struct {
	log          logging.Logger
	db           *gorm.DB
	grpcServer   *grpc.Server
	cfg          *config.Configurtion
	oryClient    oryclient.OryClient
	userapi      ports.UserAPIPort
	cerbosClient *cerbosclient.Client
}
