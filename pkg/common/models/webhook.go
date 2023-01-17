package models

import "time"

type WebhookStatus string

const (
	WebhookStatusActive      WebhookStatus = "active"
	WebhookStatusPaused      WebhookStatus = "paused"
	WebhookStatusDeactivated WebhookStatus = "deactivated"
	WebhookStatusDeleted     WebhookStatus = "deleted"
)

// https://www.solidfi.com/docs/create-a-webhook?
type WebhookDataPartial struct {
	URL         string        `json:"url"`
	Events      []string      `json:"events"`
	Description string        `json:"description"`
	Status      WebhookStatus `json:"status"`
}

type WebhookDataFull struct {
	ID        string `json:"id"`
	ProgramID string `json:"programId"`
	WebhookDataPartial
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type WebhookReceived[T any] struct {
	EventType string `json:"eventType"`
	ProgramID string `json:"programId"`
	Env       string `json:"env"`
	Data      T      `json:"data"`
}
