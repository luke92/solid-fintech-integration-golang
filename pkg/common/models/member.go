package models

import "time"

type memberDataPartial struct {
	BusinessID string `json:"businessId,omitempty"`
	// Use this for create member family
	FamilyID string `json:"familyId,omitempty"`
	PersonID string `json:"personId,omitempty"`
	// ownership percentage in the business (use whole numbers, eg. "50")
	Ownership string      `json:"ownership"`
	Title     string      `json:"title"`
	Metadata  interface{} `json:"metadata"`
}

type NewMemberData struct {
	memberDataPartial
	Person PersonDataPartial `json:"person,omitempty"`
}

type MemberDataFull struct {
	ID string `json:"id"`
	// true if this member is a control person of the business
	IsControlPerson bool `json:"isControlPerson"`
	memberDataPartial
	Person     PersonDataFull `json:"person"`
	CreatedAt  time.Time      `json:"createdAt"`
	ModifiedAt time.Time      `json:"modifiedAt"`
}
