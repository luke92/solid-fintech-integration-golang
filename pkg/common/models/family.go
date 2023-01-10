package models

import "time"

type FamilyDataPartial struct {
	Name     string      `json:"name"`
	Phone    string      `json:"phone"`
	Email    string      `json:"email"`
	Address  Address     `json:"address"`
	Metadata interface{} `json:"metadata"`
}

type FamilyStatus string

const (
	FamilityStatusActive      FamilyStatus = "active"
	FamilityStatusInactive    FamilyStatus = "inactive"
	FamilityStatusDormant     FamilyStatus = "dormant"
	FamilityStatusDeactivated FamilyStatus = "deactivated"
)

type FamilyDataFull struct {
	ID string `json:"id"`
	FamilyDataPartial
	ProgramID       string       `json:"programId"`
	CreatedPersonID string       `json:"createdPersonId"`
	Status          FamilyStatus `json:"status"`
	CreatedAt       time.Time    `json:"createdAt"`
	ModifiedAt      time.Time    `json:"modifiedAt"`
}
