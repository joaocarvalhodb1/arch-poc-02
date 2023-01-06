package app

import (
	"github.com/joaocarvalhodb1/arch-poc/services/account_service/domain"
	"github.com/joaocarvalhodb1/arch-poc/shared/contracts/protog"
	"github.com/joaocarvalhodb1/arch-poc/shared/helpers"
	"google.golang.org/grpc"
)

type AccountAppService struct {
	protog.UnimplementedAccountServiceServer
	accountRepo        domain.AccountRepository
	creditLimitService *domain.ApplyCreditLimit
	log                *helpers.Loggers
}

func NewAccountAppService(
	accountRepo domain.AccountRepository,
	creditLimitService *domain.ApplyCreditLimit,
	log *helpers.Loggers) *AccountAppService {
	return &AccountAppService{
		accountRepo:        accountRepo,
		creditLimitService: creditLimitService,
		log:                log,
	}
}

func (service *AccountAppService) RegisterServiceServer(server *grpc.Server) {
	protog.RegisterAccountServiceServer(server, service)
}
