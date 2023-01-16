package usecase

import "github.com/luke92/solid-fintech-integration-golang/pkg/types"

type IBankingService interface {
	AddPersonAndFamily(input types.AddPersonAndFamilyInput) (types.AddPersonAndFamilyOutput, error)
}
