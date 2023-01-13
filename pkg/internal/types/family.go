package types

import "github.com/luke92/solid-fintech-integration-golang/pkg/common/models"

type AddPersonAndFamilyInput = models.PersonDataPartial

type AddPersonAndFamilyOutput struct {
	PersonID string `json:"personId"`
	KYCID    string `json:"kycId"`
	IdvID    string `json:"idvId"`
	FamilyID string `json:"familyId"`
}
