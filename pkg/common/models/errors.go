package models

import (
	"encoding/json"
	"errors"
	"time"
)

type ErrorResponse struct {
	RequestID  string    `json:"requestId"`
	Method     string    `json:"method"`
	Status     int       `json:"status"`
	Code       string    `json:"code"`
	Message    string    `json:"message"`
	SysMessage string    `json:"sysMessage"`
	ProgramID  string    `json:"programId"`
	Env        string    `json:"env"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (e *ErrorResponse) ToString() string {
	out, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func (e *ErrorResponse) ToError() error {
	return errors.New(e.ToString())
}
