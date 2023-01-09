package models

import "time"

type AccountType string

const (
	AccountTypeBusiness = "businessChecking"
	AccountTypePersonal = "personalChecking"
	AccountTypeCard     = "cardAccount"
	AccountTypeClearing = "clearingAccount"
	AccountTypeFallback = "fallbackAccount"
	AccountTypeFunding  = "fundingAccount"
)

type AccountConfigCardType struct {
	Enabled bool   `json:"enabled"`
	Count   string `json:"count"`
}

type AccountConfigCards struct {
	Virtual  AccountConfigCardType `json:"virtual"`
	Physical AccountConfigCardType `json:"physical"`
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

type AccountConfigOperationConfig struct {
	Enabled bool `json:"enabled"`
	SameDay bool `json:"sameDay"`
}

type LimitsType struct {
	Receive LimitsOperations `json:"receive"`
	Send    LimitsOperations `json:"send"`
}

type LimitsAmount struct {
	Daily   string `json:"daily"`
	Monthly string `json:"monthly"`
}

type AccountConfig struct {
	Card    AccountConfigCards     `json:"card"`
	Send    AccountConfigOperation `json:"send"`
	Receive AccountConfigOperation `json:"receive"`
	Limits  LimitsType             `json:"limits"`
	//FALBACK
	//incoming
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

type AccountDataPartial struct {
	FamilyID      string      `json:"familyId"`
	Label         string      `json:"label"`
	AcceptedTerms bool        `json:"acceptedTerms"`
	Type          AccountType `json:"type"`
	Metadata      interface{} `json:"metadata"`
}

type AccountDataFull struct {
	ID                       string    `json:"id"`
	BusinessID               string    `json:"businessId"`
	RoutingNumber            string    `json:"routingNumber"`
	AccountNumber            string    `json:"accountNumber"`
	Status                   string    `json:"status"`
	ProgramID                string    `json:"programId"`
	IsVerified               bool      `json:"isVerified"`
	VerifiedAt               string    `json:"verifiedAt"`
	Interest                 string    `json:"interest"`
	Fees                     string    `json:"fees"`
	Currency                 string    `json:"currency"`
	AvailableBalance         string    `json:"availableBalance"`
	SponsorBankName          string    `json:"sponsorBankName"`
	CreatedAt                time.Time `json:"createdAt"`
	ModifiedAt               time.Time `json:"modifiedAt"`
	PendingDebit             string    `json:"pendingDebit"`
	PendingCredit            string    `json:"pendingCredit"`
	CreatedPersonID          string    `json:"createdPersonId"`
	AccountInterestFrequency string    `json:"accountInterestFrequency"`
	Config                   struct {
		Card struct {
			Virtual struct {
				Enabled bool   `json:"enabled"`
				Count   string `json:"count"`
			} `json:"virtual"`
			Physical struct {
				Enabled bool   `json:"enabled"`
				Count   string `json:"count"`
			} `json:"physical"`
		} `json:"card"`
		Send struct {
			Intrabank struct {
				Enabled bool `json:"enabled"`
			} `json:"intrabank"`
			Ach struct {
				Enabled bool `json:"enabled"`
				SameDay bool `json:"sameDay"`
			} `json:"ach"`
			Wire struct {
				Enabled bool `json:"enabled"`
			} `json:"wire"`
			Check struct {
				Enabled bool `json:"enabled"`
			} `json:"check"`
			Card struct {
				Enabled bool `json:"enabled"`
			} `json:"card"`
			DebitCard struct {
				Enabled bool `json:"enabled"`
			} `json:"debitCard"`
			InternationalWire struct {
				Enabled bool `json:"enabled"`
			} `json:"internationalWire"`
			DigitalCheck struct {
				Enabled bool `json:"enabled"`
			} `json:"digitalCheck"`
			PhysicalCard struct {
				Enabled bool `json:"enabled"`
			} `json:"physicalCard"`
			CrossBorder struct {
				Enabled bool `json:"enabled"`
			} `json:"crossBorder"`
		} `json:"send"`
		Receive struct {
			Intrabank struct {
				Enabled bool `json:"enabled"`
			} `json:"intrabank"`
			Ach struct {
				Enabled bool `json:"enabled"`
				SameDay bool `json:"sameDay"`
			} `json:"ach"`
			Check struct {
				Enabled bool `json:"enabled"`
			} `json:"check"`
			DebitCard struct {
				Enabled bool `json:"enabled"`
			} `json:"debitCard"`
		} `json:"receive"`
		Limits struct {
			Receive struct {
				Daily     string `json:"daily"`
				Monthly   string `json:"monthly"`
				Intrabank struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"intrabank"`
				Ach struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"ach"`
				Check struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"check"`
				DebitCard struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"debitCard"`
			} `json:"receive"`
			Send struct {
				Daily     string `json:"daily"`
				Monthly   string `json:"monthly"`
				Intrabank struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"intrabank"`
				Ach struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"ach"`
				DomesticWire struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"domesticWire"`
				InternationalWire struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"internationalWire"`
				Check struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"check"`
				Card struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"card"`
				DebitCard struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"debitCard"`
				CrossBorder struct {
					Daily   string `json:"daily"`
					Monthly string `json:"monthly"`
				} `json:"crossBorder"`
			} `json:"send"`
		} `json:"limits"`
		Fallback struct {
			Received interface{} `json:"received"`
		} `json:"fallback"`
		Incoming struct {
			AchPull struct {
				Enabled            bool        `json:"enabled"`
				AllowedOriginators interface{} `json:"allowedOriginators"`
				BlockedOriginators interface{} `json:"blockedOriginators"`
			} `json:"achPull"`
		} `json:"incoming"`
	} `json:"config"`
}
