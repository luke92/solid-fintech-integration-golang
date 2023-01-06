package solid

import (
	"encoding/json"
	"fmt"

	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

func (service *SolidService) AddPerson(model models.PersonDataPartial) (models.PersonDataFull, error) {
	relativeURL := "/v1/person"
	method := "POST"
	payload, err := json.Marshal(model)
	if err != nil {
		fmt.Println("Error Add Person Solid - Parse Request", err)
		return models.PersonDataFull{}, err
	}

	input := sendRequestInput{
		Method:      method,
		RelativeURL: relativeURL,
		Payload:     payload,
	}

	output, err := service.sendRequest(input)
	if err != nil {
		fmt.Println("Error Add Person Solid - Send request", err)
		return models.PersonDataFull{}, err
	}

	if output.ErrorSolid != nil {
		return models.PersonDataFull{}, output.ErrorSolid.ToError()
	}

	var response models.PersonDataFull

	err = json.Unmarshal(output.Data, &response)
	if err != nil {
		fmt.Println("Error Add Person Solid - Parse Response")
		return models.PersonDataFull{}, err
	}

	return response, nil
}

func (service *SolidService) SubmitKYC(personID string) (models.KYC, error) {
	errorctx := "submit-kyc-error"
	var response models.KYC

	relativeURL := fmt.Sprintf("/v1/person/%s/kyc", personID)
	method := "POST"

	input := sendRequestInput{
		Method:      method,
		RelativeURL: relativeURL,
		PersonId:    personID,
	}

	output, err := service.sendRequest(input)
	if err != nil {
		fmt.Println(errorctx+"-send-request", err)
		return response, err
	}

	if output.ErrorSolid != nil {
		return response, output.ErrorSolid.ToError()
	}

	err = json.Unmarshal(output.Data, &response)
	if err != nil {
		fmt.Println(errorctx + "-parse-response")
		return response, err
	}

	return response, nil
}

func (service *SolidService) SubmitIDV(personID string) (models.IDV, error) {
	errorctx := "submit-idv-error"
	var response models.IDV

	relativeURL := fmt.Sprintf("/v1/person/%s/idv", personID)
	method := "POST"

	input := sendRequestInput{
		Method:      method,
		RelativeURL: relativeURL,
		PersonId:    personID,
	}

	output, err := service.sendRequest(input)
	if err != nil {
		fmt.Println(errorctx+"-send-request", err)
		return response, err
	}

	if output.ErrorSolid != nil {
		return response, output.ErrorSolid.ToError()
	}

	err = json.Unmarshal(output.Data, &response)
	if err != nil {
		fmt.Println(errorctx + "-parse-response")
		return response, err
	}

	return response, nil
}
