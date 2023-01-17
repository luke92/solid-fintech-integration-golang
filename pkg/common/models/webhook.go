package models

import "time"

/*
Below you will find links to webhook samples sent to webhook listeners. Event types and descriptions:

account.created - when a new account is created

account.credit - when an account records a credit transaction

account.debit - when an account records a debit transaction

account.updated - when account information is updated

ach.auth.request - when a Real Time Authorization (RTA) ACH request occurs

ach.originated.credit - on origination of an ACH credit transactions

ach.originated.debit - on origination of an ACH debit transactions

ach.originated.updated - when an ACH transaction is updated

ach.received.credit - when an ACH credit transaction is received

ach.received.debit - when an ACH debit transaction is received

ach.received.declined - when an ACH received transaction is declined (eg. insufficient balance or blocked account)

ach.reversal.credit - when an ACH is reversed resulting in a credit

ach.reversal.debit - when an ACH is reversed resulting in a debit

business.created - when a new business is created

business.updated - when business information Is updated

business.kyb.approved - when a business KYB is approved

business.kyb.declined - when a business KYB is declined

business.kyb.review - when a business KYB goes into review

card.auth.request - when a Real Time Authorization (RTA) card auth request occurs

card.canceled - when a card is canceled

card.declined - when a card transaction is declined

card.deposit - when cash is deposited at an ATM on a debit card

card.enriched - when a card spend is updated with merchant enriched data

card.refund - when a card refund transaction occurs

card.replaced - when a card is replaced

card.reversal - when a card transaction is reversed

card.reward.credit - when a Dosh card transaction reward credit is credited

card.reward.pending - when a Dosh card transaction reward is potentially earned

card.reward.canceled - when a Dosh card transaction reward is canceled

card.spend - when a card authorization occurs

card.spend.update - when a card authorization is updated

card.updated - when card information is updated

card.wallet.created - when card is added to a digital wallet

card.wallet.updated - when a digital wallet is updated

check.originated.credit - when a check is deposited to an account

check.originated.debit - when a check is sent, debiting the account

check.reversal.credit - when a sent check is reversed

check.reversal.debit - when a check deposit is reversed

contact.created - when a contact is created for an account

contact.updated - when contact information is updated

family.created - when a new family is created

family.updated - when family information updated

intrabank.originated.credit - when an intrabank pull is originated

intrabank.originated.debit - when an intrabank push is originated

intrabank.received.credit - when an intrabank push is received

intrabank.received.debit - when an intrabank pull is received

member.created - when a business/family member is updated

member.updated - when a business/family member is updated

person.created - when a new person is created

person.kyc.approved - when a person KYC is approved

person.kyc.declined - when a person KYC is declined

person.kyc.review - when a person KYC goes into review

person.kyc.updated - when a person KYC is updated

person.updated - when a person's information Is updated

solidcard.canceled - when a Send a Card (Solid card) is canceled

solidcard.delivered - when a Send a Card (Solid card) is delivered

solidcard.originated.debit - when a Send a Card (Solid card) transaction is originated

solidcard.reversal.credit - when a Send a Card (Solid card) transaction is reversed (when card is canceled)

solidcard.topup - when a Send a Card (Solid card) is topped up

transfer.created - when a transfer is created

transfer.updated - when a transfer is updated

wallet.created - when a crypto wallet is created

wallet.credit - when a crypto wallet is credited

wallet.debit - when a crypto wallet is debited

wallet.transfer.created - when a crypto transfer is created

wallet.transfer.updated - when a crypto transfer is updated

wallet.updated - when a crypto wallet's information is updated

‍wire.received.credit - when an incoming wire is received

wire.originated.debit - when an outgoing wire is orginated

wire.originated.updated - when an outgoing wire is updated (eg. IMAD/OMAD added)

wire.reversal.debit - when an incoming wire is reversed

wire.reversal.credit - when an outgoing wire is reversed
*/
type WebhookEvent string

type WebhookStatus string

const (
	WebhookStatusActive      WebhookStatus = "active"
	WebhookStatusPaused      WebhookStatus = "paused"
	WebhookStatusDeactivated WebhookStatus = "deactivated"
	WebhookStatusDeleted     WebhookStatus = "deleted"
)

// https://www.solidfi.com/docs/create-a-webhook?
type WebhookDataPartial struct {
	URL         string         `json:"url"`
	Events      []WebhookEvent `json:"events"`
	Description string         `json:"description"`
	Status      WebhookStatus  `json:"status"`
}

type WebhookDataFull struct {
	ID        string `json:"id"`
	ProgramID string `json:"programId"`
	WebhookDataPartial
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type WebhookReceived[T any] struct {
	EventType WebhookEvent `json:"eventType"`
	ProgramID string       `json:"programId"`
	Env       string       `json:"env"`
	Data      T            `json:"data"`
}
