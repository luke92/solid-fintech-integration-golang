package banking

import (
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
	"github.com/luke92/solid-fintech-integration-golang/pkg/usecase"
)

type BankingService struct {
	service solid.Service
}

func NewBankingService(service solid.Service) usecase.IBankingService {
	return &BankingService{
		service: service,
	}
}
