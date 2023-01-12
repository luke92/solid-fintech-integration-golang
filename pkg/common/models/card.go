package models

import "time"

type CardType string

const (
	CardTypeVirtual  CardType = "virtual"
	CardTypePhysical CardType = "physical"
)

type LimitInterval string

const (
	LimitIntervalDaily          LimitInterval = "daily"
	LimitIntervalWeekly         LimitInterval = "weekly"
	LimitIntervalMonthly        LimitInterval = "monthly"
	LimitIntervalYearly         LimitInterval = "yearly"
	LimitIntervalAllTime        LimitInterval = "allTime"
	LimitIntervalPerTransaction LimitInterval = "perTransaction"
)

type ShippingNewCard struct {
	ShippingAddress Address `json:"shippingAddress"`
}

type CardBin string

const (
	CardBinDebit  CardBin = "debit"
	CardBinCredit CardBin = "credit"
)

type cardDataPartial struct {
	// unique id of the bank account to which the card belongs
	AccountID string   `json:"accountId"`
	CardType  CardType `json:"cardType"`
	Label     string   `json:"label"`
	// 	spend limit amount of the card ("limitAmount": "120.00")
	LimitAmount   string        `json:"limitAmount"`
	LimitInterval LimitInterval `json:"limitInterval"`
	// USD
	Currency string  `json:"currency"`
	Bin      CardBin `json:"bin"`
	// name of the person to be embossed on the card, max 23 characters, not allowing: ! " & ' * + , : ; < = > @ _
	// Not Necessary for Virtual Cards
	EmbossingPerson string `json:"embossingPerson,omitempty"`
	// name of the person to be embossed on the card, max 23 characters, not allowing: ! " & ' * + , : ; < = > @ _
	// Not Necessary for Virtual Cards
	EmbossingBusiness string `json:"embossingBusiness,omitempty"`
}

type NewCard struct {
	cardDataPartial
	BillingAddress Address `json:"billingAddress"`
	// required only for a physical card
	Shipping ShippingNewCard `json:"shipping,omitempty"`
	Metadata interface{}     `json:"metadata,omitempty"`
}

type Cardholder struct {
	ID             string    `json:"id"`
	PersonID       string    `json:"personId"`
	BillingAddress Address   `json:"billingAddress"`
	CreatedAt      time.Time `json:"createdAt"`
	ModifiedAt     time.Time `json:"modifiedAt"`
	Name           string    `json:"name"`
}

type DeliveryStatus string

const (
	DeliveryStatusPending   DeliveryStatus = "pending"
	DeliveryStatusShipped   DeliveryStatus = "shipped"
	DeliveryStatusDelivered DeliveryStatus = "delivered"
	DeliveryStatusReturned  DeliveryStatus = "returned"
	DeliveryStatusFailure   DeliveryStatus = "failure"
	DeliveryStatusCanceled  DeliveryStatus = "canceled"
)

type Shipping struct {
	ID string `json:"id"`
	ShippingNewCard
	// 	estimated time of arrival for the shipped card
	Eta            string         `json:"eta"`
	DeliveryStatus DeliveryStatus `json:"deliveryStatus"`
	CreatedAt      time.Time      `json:"createdAt"`
	ModifiedAt     time.Time      `json:"modifiedAt"`
	// delivery partner (eg. USPS) used to ship the card
	DeliveryPartner string `json:"deliveryPartner"`
	// tracking number if USPS Priority or UPS Next Day
	TrackingNumber string `json:"trackingNumber"`
	// tracking URL if USPS Priority or UPS Next Day
	TrackingURL string `json:"trackingUrl"`
}

type CardStatus string

const (
	CardStatusActive            CardStatus = "active"
	CardStatusInactive          CardStatus = "inactive" // (use to temporarily freeze)
	CardStatusCanceled          CardStatus = "canceled" // (use to permanently cancel)
	CardStatusPendingActivation CardStatus = "pendingActivation"
)

type ExpireCardData struct {
	// Month with 2 digits (Example: 01)
	ExpiryMonth string `json:"expiryMonth"`
	// Year with 4 digits (Example: 2024)
	ExpiryYear string `json:"expiryYear"`
}

// By default, a physical card is created as inactive. Use the Activate a Card endpoint to set its status to active. A virtual card is active on creation, hence it does not need to be activated.
type ExpireAndLast4CardData struct {
	ExpireCardData
	// Last 4 digits of the card number
	Last4 string `json:"last4"`
}

type CardDataFull struct {
	ID string `json:"id"`
	cardDataPartial
	BusinessID string     `json:"businessId"`
	ProgramID  string     `json:"programId"`
	Cardholder Cardholder `json:"cardholder"`
	Shipping   Shipping   `json:"shipping"`
	ExpireAndLast4CardData
	CardStatus CardStatus `json:"cardStatus"`
	// A virtual card is active on creation, Physical will have ActivatedAt after Activate Card
	ActivatedAt     TimeWrapper `json:"activatedAt,omitempty"`
	CreatedAt       time.Time   `json:"createdAt"`
	ModifiedAt      time.Time   `json:"modifiedAt"`
	CreatedPersonID string      `json:"createdPersonId"`
	// an array of mcc codes to allow on the card. Example ["0742","0763"] https://github.com/greggles/mcc-codes/blob/main/mcc_codes.json
	AllowedCategories []string `json:"allowedCategories"`
	AvailableLimit    string   `json:"availableLimit"`
	// an array of mcc codes to block on the card. Example ["0742","0763"] https://github.com/greggles/mcc-codes/blob/main/mcc_codes.json
	BlockedCategories []string `json:"blockedCategories"`
	// card theme, if the program has multiple card themes
	Theme string `json:"theme"`
	// atm access for the card, true means enabled (by default), false means disabled
	AtmAccess bool `json:"atmAccess"`
	// an array of merchant names to allow on the card. Example ["Subway","Campus Cafe"]
	AllowedMerchants []string `json:"allowedMerchants"`
	// 	an array of merchant names to allow on the card. Example ["Arco","Valero"]
	BlockedMerchants []string `json:"blockedMerchants,omitempty"`
	WalletID         string   `json:"walletId"`
	FamilyID         string   `json:"familyId"`
}

type ActivateCardResponse struct {
	ID     string     `json:"id"`
	Status CardStatus `json:"status"`
}

type CreateCardPinTokenResponse struct {
	ID       string `json:"id"`
	PinToken string `json:"pinToken"`
}

type CardSetNewPinRequest struct {
	ExpireAndLast4CardData
	Pin string `json:"pin"`
}

type CardSetNewPinResponse struct {
	ID string `json:"id"`
}

type CreateCardShowTokenResponse struct {
	ID        string `json:"id"`
	ShowToken string `json:"showToken"`
}

type CardSecurityData struct {
	ID         string `json:"id"`
	CardNumber string `json:"cardNumber"`
	Cvv        string `json:"cvv"`
	ExpireCardData
}
