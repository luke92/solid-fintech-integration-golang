package models

import "time"

type AccountType string

const (
	AccountTypeBusiness AccountType = "businessChecking"
	AccountTypePersonal AccountType = "personalChecking" // Use this for Family Account
	AccountTypeCard     AccountType = "cardAccount"
	AccountTypeClearing AccountType = "clearingAccount"
	AccountTypeFallback AccountType = "fallbackAccount"
	AccountTypeFunding  AccountType = "fundingAccount"
)

type AccountConfigCardType struct {
	Enabled bool   `json:"enabled"`
	Count   string `json:"count"`
}

type AccountConfigCards struct {
	Virtual  AccountConfigCardType `json:"virtual"`
	Physical AccountConfigCardType `json:"physical"`
}

type AccountConfigOperationConfig struct {
	Enabled bool `json:"enabled"`
	SameDay bool `json:"sameDay"`
}

type AccountConfigOperation struct {
	Intrabank         AccountConfigOperationConfig `json:"intrabank"`
	Ach               AccountConfigOperationConfig `json:"ach"`
	Wire              AccountConfigOperationConfig `json:"wire"`
	Check             AccountConfigOperationConfig `json:"check"`
	Card              AccountConfigOperationConfig `json:"card"`
	DebitCard         AccountConfigOperationConfig `json:"debitCard"`
	InternationalWire AccountConfigOperationConfig `json:"internationalWire"`
	DigitalCheck      AccountConfigOperationConfig `json:"digitalCheck"`
	PhysicalCard      AccountConfigOperationConfig `json:"physicalCard"`
	CrossBorder       AccountConfigOperationConfig `json:"crossBorder"`
}

type LimitsAmount struct {
	// Limit amount (Example json attribute "daily": "1000.00")
	Daily   string `json:"daily"`
	Monthly string `json:"monthly"`
}

type LimitsOperations struct {
	LimitsAmount
	Intrabank         LimitsAmount `json:"intrabank"`
	Ach               LimitsAmount `json:"ach"`
	DomesticWire      LimitsAmount `json:"domesticWire"`
	InternationalWire LimitsAmount `json:"internationalWire"`
	Check             LimitsAmount `json:"check"`
	Card              LimitsAmount `json:"card"`
	DebitCard         LimitsAmount `json:"debitCard"`
	CrossBorder       LimitsAmount `json:"crossBorder"`
}

type LimitsType struct {
	Receive LimitsOperations `json:"receive"`
	Send    LimitsOperations `json:"send"`
}

type Fallback struct {
	Received interface{} `json:"received"`
}

type AchPull struct {
	Enabled            bool        `json:"enabled"`
	AllowedOriginators interface{} `json:"allowedOriginators"`
	BlockedOriginators interface{} `json:"blockedOriginators"`
}

type Incoming struct {
	AchPull AchPull `json:"achPull"`
}

type AccountConfig struct {
	Card     AccountConfigCards     `json:"card"`
	Send     AccountConfigOperation `json:"send"`
	Receive  AccountConfigOperation `json:"receive"`
	Limits   LimitsType             `json:"limits"`
	Fallback Fallback               `json:"fallback"`
	Incoming Incoming               `json:"incoming"`
}

type AccountDataPartial struct {
	FamilyID      string      `json:"familyId"`
	Label         string      `json:"label"`
	AcceptedTerms bool        `json:"acceptedTerms"`
	Type          AccountType `json:"type"`
	Metadata      interface{} `json:"metadata"`
}

type AccountStatus string

const (
	AccountStatusActive        AccountStatus = "active"
	AccountStatusDebitBlocked  AccountStatus = "debitBlocked"  // (debits frozen)
	AccountStatusCreditBlocked AccountStatus = "creditBlocked" // (credits frozen)
	AccountStatusBlocked       AccountStatus = "blocked"       // (both frozen)
	AccountStatusClosed        AccountStatus = "closed"        // (permanently closed)
)

type AccountDataFull struct {
	AccountDataPartial
	ID         string `json:"id"`
	BusinessID string `json:"businessId"`
	// 	9 digit routing number of the bank account
	RoutingNumber string `json:"routingNumber"`
	// 	16 digit account number of the bank account
	AccountNumber string        `json:"accountNumber"`
	Status        AccountStatus `json:"status"`
	ProgramID     string        `json:"programId"`
	IsVerified    bool          `json:"isVerified"`
	VerifiedAt    string        `json:"verifiedAt"`
	// 	interest percentage if an interest earning account
	Interest string `json:"interest"`
	// fees charged per month if it is a fee based account
	Fees string `json:"fees"`
	// 	currency of the bank account, currently supported: - USD
	Currency string `json:"currency"`
	// available balance in the bank account
	AvailableBalance string    `json:"availableBalance"`
	SponsorBankName  string    `json:"sponsorBankName"`
	CreatedAt        time.Time `json:"createdAt"`
	ModifiedAt       time.Time `json:"modifiedAt"`
	// pending debits (money going out) from the bank account
	PendingDebit string `json:"pendingDebit"`
	// pending credits (money coming in) in to the bank account
	PendingCredit   string `json:"pendingCredit"`
	CreatedPersonID string `json:"createdPersonId"`
	// frequency at which interest is paid out (if any)
	AccountInterestFrequency string        `json:"accountInterestFrequency"`
	Config                   AccountConfig `json:"config"`
}
