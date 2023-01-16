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
	GetAccounts(personID string) (models.List[models.AccountDataFull], error)

	// Member

	AddMember(personID string, model models.NewMemberData) (models.MemberDataFull, error)

	// CARD

	AddCard(personID string, model models.NewCard) (models.CardDataFull, error)
	ActivateCard(personID string, cardID string, model models.ExpireAndLast4CardData) (models.ActivateCardResponse, error)
	CreatePinTokenCard(personID string, cardID string) (models.CreateCardPinTokenResponse, error)
	CreateShowTokenCard(personID string, cardID string) (models.CreateCardShowTokenResponse, error)

	// CONTACT

	AddContact(personID string, model models.ContactDataPartial) (models.ContactDataFull, error)
	GetContacts(personID string, accountID string) (models.List[models.ContactDataFull], error)
	GetBankInfo(personID string, accountType models.AccountBankType, routingNumber string) (models.BankInfo, error)
}
