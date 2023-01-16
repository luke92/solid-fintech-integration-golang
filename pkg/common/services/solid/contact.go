package solid

import (
	"encoding/json"
	"fmt"

	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

func (service *SolidService) AddContact(personID string, model models.ContactDataPartial) (models.ContactDataFull, error) {
	var response models.ContactDataFull
	ctxError := "add-contact"

	relativeURL := "/v1/contact"
	method := "POST"
	payload, err := json.Marshal(model)
	if err != nil {
		fmt.Println(ctxError+"-parse-request", err)
		return response, err
	}

	input := sendRequestInput{
		Method:      method,
		RelativeURL: relativeURL,
		Payload:     payload,
		PersonId:    personID,
	}

	output, err := service.sendRequest(input)
	if err != nil {
		fmt.Println(ctxError+"-send-request", err)
		return response, err
	}

	if output.ErrorSolid != nil {
		return response, output.ErrorSolid.ToError()
	}

	err = json.Unmarshal(output.Data, &response)
	if err != nil {
		fmt.Println(ctxError + "-parse-response")
		return response, err
	}

	return response, nil
}

func (service *SolidService) GetContacts(personID string, accountID string) (models.List[models.ContactDataFull], error) {
	var response models.List[models.ContactDataFull]
	ctxError := "get-contacts"

	relativeURL := fmt.Sprintf("/v1/contact?accountId=%s", accountID)
	method := "GET"

	input := sendRequestInput{
		Method:      method,
		RelativeURL: relativeURL,
		PersonId:    personID,
	}

	output, err := service.sendRequest(input)
	if err != nil {
		fmt.Println(ctxError+"-send-request", err)
		return response, err
	}

	if output.ErrorSolid != nil {
		return response, output.ErrorSolid.ToError()
	}

	err = json.Unmarshal(output.Data, &response)
	if err != nil {
		fmt.Println(ctxError + "-parse-response")
		return response, err
	}

	return response, nil
}

func (service *SolidService) GetBankInfo(personID string, accountType models.AccountBankType, routingNumber string) (models.BankInfo, error) {
	var response models.BankInfo
	ctxError := "get-bank-info"

	relativeURL := fmt.Sprintf("/v1/contact/bank?type=%s&routingNumber=%s", accountType, routingNumber)
	method := "GET"

	input := sendRequestInput{
		Method:      method,
		RelativeURL: relativeURL,
		PersonId:    personID,
	}

	output, err := service.sendRequest(input)
	if err != nil {
		fmt.Println(ctxError+"-send-request", err)
		return response, err
	}

	if output.ErrorSolid != nil {
		return response, output.ErrorSolid.ToError()
	}

	err = json.Unmarshal(output.Data, &response)
	if err != nil {
		fmt.Println(ctxError + "-parse-response")
		return response, err
	}

	return response, nil
}
