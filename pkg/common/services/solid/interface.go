package solid

import "github.com/luke92/solid-fintech-integration-golang/pkg/common/models"

type Service interface {
	// Person

	AddPerson(model models.PersonDataPartial) (models.PersonDataFull, error)
	SubmitKYC(personID string) (models.KYC, error)
	SubmitIDV(personID string) (models.IDV, error)

	// Family

	AddFamily(personID string, model models.FamilyDataPartial) (models.FamilyDataFull, error)

	// Account

	AddAccount(personID string, model models.AccountDataPartial) (models.AccountDataFull, error)
}
