package solid

import (
	"encoding/json"
	"fmt"

	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

func (service *SolidService) AddAccount(personID string, model models.AccountDataPartial) (models.AccountDataFull, error) {
	var response models.AccountDataFull
	ctxError := "add-account"

	relativeURL := "/v1/account"
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

func (service *SolidService) GetAccounts(personID string) (models.List[models.AccountDataFull], error) {
	var response models.List[models.AccountDataFull]
	ctxError := "get-accounts"

	relativeURL := fmt.Sprintf("/v1/account?personId=%s", personID)
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
