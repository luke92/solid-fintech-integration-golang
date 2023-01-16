package models

import "time"

// Used for Send or Receive Intrabank
type FundIntrabankRequest struct {
	AccountID string `json:"accountId"`
	ContactID string `json:"contactId"`
	// "amount": "500.00",
	Amount      string      `json:"amount"`
	Description string      `json:"description"`
	Metadata    interface{} `json:"metadata"`
}

type FundAchType string

const (
	AchTypeSameDay FundAchType = "sameDay"
	AchTypeNextDay FundAchType = "nextDay"
	AchTypeNone    FundAchType = ""
)

// Used for Send or Receive Ach
type FundAchRequest struct {
	FundIntrabankRequest
	// You can also pass the type for the ACH (sameDay, nextDay). Default is nextDay
	Type FundAchType `json:"type"`
}

type FundAccountType string

const (
	FundAccountTypeBusinessChecking FundAccountType = "businessChecking"
	FundAccountTypePersonalChecking FundAccountType = "personalChecking"
	FundAccountTypeBusinessSavings  FundAccountType = "businessSavings"
	FundAccountTypePersonalSavings  FundAccountType = "personalSavings"
	FundAccountTypeNone             FundAccountType = ""
)

type FundStatus string

const (
	FundStatusCompleted FundStatus = "completed"
	FundStatusPending   FundStatus = "pending"
	FundStatusSettled   FundStatus = "settled"
	FundStatusCanceled  FundStatus = "completed"
	FundStatusRefund    FundStatus = "refund"
	FundStatusDeclined  FundStatus = "declined"
)

type FundTransferSubType string

const (
	FundTransferSubTypeOriginated            FundTransferSubType = "originated"
	FundTransferSubTypeReceived              FundTransferSubType = "received"
	FundTransferSubTypeAuth                  FundTransferSubType = "auth"
	FundTransferSubTypeAuthUpdate            FundTransferSubType = "authUpdate"
	FundTransferSubTypeReversal              FundTransferSubType = "reversal"
	FundTransferSubTypeRefund                FundTransferSubType = "refund"
	FundTransferSubTypeInterest              FundTransferSubType = "interest"
	FundTransferSubTypeFees                  FundTransferSubType = "fees"
	FundTransferSubTypePenalty               FundTransferSubType = "penalty"
	FundTransferSubTypePromotionalCredit     FundTransferSubType = "promotionalCredit"
	FundTransferSubTypeAccountCreationCredit FundTransferSubType = "accountCreationCredit"
)

type Fund struct {
	ID       string `json:"id"`
	BankName string `json:"bankName"`
	FundAchRequest
	// 	name of the contact, either full name of person or legal name of business
	Name             string              `json:"name"`
	AccountNumber    string              `json:"accountNumber"`
	RoutingNumber    string              `json:"routingNumber"`
	Address          Address             `json:"address"`
	Status           FundStatus          `json:"status"`
	TxnType          TransactionType     `json:"txnType"`
	TransferType     TransferType        `json:"transferType"`
	TransferSubType  FundTransferSubType `json:"transferSubType"`
	CreatedAt        time.Time           `json:"createdAt"`
	ModifiedAt       time.Time           `json:"modifiedAt"`
	TransferredAt    time.Time           `json:"transferredAt"`
	AccountType      FundAccountType     `json:"accountType"`
	Iban             string              `json:"iban"`
	Valid            string              `json:"valid"`
	ParentTransferID string              `json:"parentTransferId"`
	// https://global-uploads.webflow.com/616c7256b52c543a920623bd/62bd0bf5ea57f8ea867aab30_Solid%20Transfer%20Review%20Code%20Matrix.pdf
	ReviewCode          string      `json:"reviewCode"`
	ReviewMessage       string      `json:"reviewMessage"`
	Title               string      `json:"title"`
	Card                interface{} `json:"card"`
	DestinationAmount   string      `json:"destinationAmount"`
	DestinationCurrency string      `json:"destinationCurrency"`
	PurposeCode         string      `json:"purposeCode"`
	Bin                 string      `json:"bin"`
}
