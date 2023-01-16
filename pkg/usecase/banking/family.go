package banking

import (
	"fmt"

	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
	"github.com/luke92/solid-fintech-integration-golang/pkg/types"
)

func (b *BankingService) AddPersonAndFamily(input types.AddPersonAndFamilyInput) (types.AddPersonAndFamilyOutput, error) {
	var output types.AddPersonAndFamilyOutput

	ctx := "add-person-and-family"
	person, err := b.service.AddPerson(input)
	if err != nil {
		fmt.Println(ctx+"-service-add-person", err)
		return output, err
	}

	output.PersonID = person.ID

	kyc, err := b.service.SubmitKYC(output.PersonID)
	if err != nil {
		fmt.Println(ctx+"-service-submit-kyc", err)
		return output, err
	}

	output.KYCID = kyc.ID

	idv, err := b.service.SubmitIDV(output.PersonID)
	if err != nil {
		fmt.Println(ctx+"-service-submit-idv", err)
		return output, err
	}

	output.IdvID = idv.ID

	newFamily := models.FamilyDataPartial{
		Name:    input.LastName,
		Phone:   input.Phone,
		Email:   input.Email,
		Address: input.Address,
	}
	family, err := b.service.AddFamily(output.PersonID, newFamily)
	if err != nil {
		fmt.Println(ctx+"-service-add-family", err)
		return output, err
	}

	output.FamilyID = family.ID
	return output, nil

}
