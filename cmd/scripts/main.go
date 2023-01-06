package main

import (
	"flag"
)

var (
	BaseURL    = flag.String("base-url", "https://test-api.solidfi.com/v1/", "BASE URL Solid Fintech")
	SdApiKey   = flag.String("sd-api-key", "", "API KEY of Solid Fintech")
	PersonID   = flag.String("personId", "", "Person ID")
	BusinessID = flag.String("businessId", "", "Business ID")
	FamilyID   = flag.String("familyId", "", "Family ID")
	MemberID   = flag.String("memberId", "", "Member ID")
	AccountID  = flag.String("accountId", "", "Account ID")
	ContactID  = flag.String("contactId", "", "Contact ID")
	TransferID = flag.String("transferId", "", "Transfer ID")
	CardID     = flag.String("cardId", "", "Card ID")
	WalletID   = flag.String("walletId", "", "Wallet ID")
	WebhookID  = flag.String("webhookId", "", "Webhook ID")
	ProgramID  = flag.String("programId", "", "Program ID")
	TicketID   = flag.String("ticketId", "", "Ticket ID") // Help Desk
)

func main() {
}
