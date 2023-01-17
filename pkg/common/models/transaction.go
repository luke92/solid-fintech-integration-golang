package models

import "time"

type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "debit"
	TransactionTypeCredit TransactionType = "credit"
)

type TransferType string

const (
	// account related transfer types

	TransferTypeIntrabank         TransferType = "intrabank"
	TransferTypeAch               TransferType = "ach"
	TransferTypeCard              TransferType = "card"
	TransferTypeDomesticWire      TransferType = "domesticWire"
	TransferTypeInternationalWire TransferType = "internationalWire"
	TransferTypeDebitCard         TransferType = "debitCard"
	TransferTypeSolidCard         TransferType = "solidCard"
	TransferTypePhysicalCheck     TransferType = "physicalCheck"
	// CRIPTO

	TransferTypeBuy     TransferType = "buy"
	TransferTypeSell    TransferType = "sell"
	TransferTypeSend    TransferType = "send"
	TransferTypeReceive TransferType = "receive"
)

type TransferSubType string

const (
	// originated = transfer is initiated
	TransferSubTypeOriginated TransferSubType = "originated"
	// auth = card authorization
	TransferSubTypeAuth TransferSubType = "auth"
	// authUpdate = card authorization update
	TransferSubTypeAuthUpdate TransferSubType = "authUpdate"
	// reversal = transaction is reversed
	TransferSubTypeReversal TransferSubType = "reversal"
	// refund = card transaction is refunded (may be less than auth amount)
	TransferSubTypeRefund TransferSubType = "refund"
	// accountCreationCredit = $0 transaction on account opening
	TransferSubTypeAccountCreationCredit TransferSubType = "accountCreationCredit"
	// received = user receives a transfer
	TransferSubTypeReceived TransferSubType = "received"
	// accountClosing = transaction while moving funds for closing account
	TransferSubTypeAccountClosing TransferSubType = "accountClosing"
	// pull = incoming debit pull
	TransferSubTypePull TransferSubType = "push"
	// push = incoming debit push
	TransferSubTypePush TransferSubType = "pull"
	// overdraft = transfer of funds to/from a fallback account
	TransferSubTypeOverdraft TransferSubType = "overdraft"
	// atm = ATM transaction
	TransferSubTypeAtm TransferSubType = "atm"
	// topup = funds added to Send a Card -> https://www.solidfi.com/docs/send-a-card?
	TransferSubTypeTopup TransferSubType = "topup"
	// reward = merchant funded reward on card spend -> Dosh https://www.solidfi.com/docs/merchant-funded-rewards?
	TransferSubTypeReward TransferSubType = "reward"
)

type TransferStatus string

const (
	// pending = in review or waiting for approval
	TransferStatusPending  TransferStatus = "pending"
	TransferStatusCanceled TransferStatus = "canceled"
	TransferStatusSettled  TransferStatus = "settled"
)

type TransactionAch struct {
	TransactionIntrabank
	Type        AchType `json:"type"`
	TraceNumber string  `json:"traceNumber"`
	AuthID      string  `json:"authId"`
}

type TransactionIntrabank struct {
	TransferID    string `json:"transferId"`
	ContactID     string `json:"contactId"`
	Name          string `json:"name"`
	RoutingNumber string `json:"routingNumber"`
	BankName      string `json:"bankName"`
}

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Merchant struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Website  string   `json:"website"`
	Logo     string   `json:"logo"`
	Address  Address  `json:"address"`
	Location Location `json:"location"`
}

type EnrichedData struct {
	Merchant        Merchant      `json:"merchant"`
	ChartOfAccounts []interface{} `json:"chartOfAccounts"`
	Labels          []interface{} `json:"labels"`
}

type Transaction struct {
	ID           string          `json:"id"`
	TxnType      TransactionType `json:"txnType"`
	Title        string          `json:"title"`
	Amount       string          `json:"amount"`
	TransferType TransferType    `json:"transferType"`
	SubType      TransferSubType `json:"subType"`
	Description  string          `json:"description"`
	Balance      string          `json:"balance"`
	// https://www.solidfi.com/docs/transaction-status-codes?
	Status           TransferStatus `json:"status"`
	AccountID        string         `json:"accountId"`
	BusinessID       string         `json:"businessId"`
	ProgramID        string         `json:"programId"`
	PersonID         string         `json:"personId"`
	FamilyID         string         `json:"familyId"`
	ParentTransferID string         `json:"parentTransferId"`
	ParentTxnID      string         `json:"parentTxnId"`
	// https://global-uploads.webflow.com/616c7256b52c543a920623bd/62bd0bf5ea57f8ea867aab30_Solid%20Transfer%20Review%20Code%20Matrix.pdf
	ReviewCode        string                `json:"reviewCode"`
	ReviewMessage     string                `json:"reviewMessage"`
	Intrabank         *TransactionIntrabank `json:"intrabank"`
	Ach               TransactionAch        `json:"ach"`
	Card              interface{}           `json:"card"`
	DomesticWire      interface{}           `json:"domesticWire"`
	InternationalWire interface{}           `json:"internationalWire"`
	DigitalCheck      interface{}           `json:"digitalCheck"`
	PhysicalCheck     interface{}           `json:"physicalCheck"`
	DebitCard         interface{}           `json:"debitCard"`
	SolidCard         interface{}           `json:"solidCard"`
	CrossBorder       interface{}           `json:"crossBorder"`
	Buy               interface{}           `json:"buy"`
	Sell              interface{}           `json:"sell"`
	EnrichedData      EnrichedData          `json:"enrichedData"`
	CreatedAt         time.Time             `json:"createdAt"`
	ModifiedAt        time.Time             `json:"modifiedAt"`
	TxnDate           time.Time             `json:"txnDate"`
	Metadata          interface{}           `json:"metadata"`
}
