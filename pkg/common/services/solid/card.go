package solid

import (
	"encoding/json"
	"fmt"

	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

func (service *SolidService) AddCard(personID string, model models.NewCard) (models.CardDataFull, error) {
	var response models.CardDataFull
	ctxError := "add-card"

	relativeURL := "/v1/card"
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

func (service *SolidService) ActivateCard(personID string, cardID string, model models.ExpireAndLast4CardData) (models.ActivateCardResponse, error) {
	var response models.ActivateCardResponse
	ctxError := "activate-card"

	relativeURL := fmt.Sprintf("/v1/card/%s/activate", cardID)
	method := "PATCH"
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

func (service *SolidService) CreatePinTokenCard(personID string, cardID string) (models.CreateCardPinTokenResponse, error) {
	var response models.CreateCardPinTokenResponse
	ctxError := "create-pin-token-card"

	relativeURL := fmt.Sprintf("/v1/card/%s/pintoken", cardID)
	method := "POST"

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

func (service *SolidService) CreateShowTokenCard(personID string, cardID string) (models.CreateCardShowTokenResponse, error) {
	var response models.CreateCardShowTokenResponse
	ctxError := "create-show-token-card"

	relativeURL := fmt.Sprintf("/v1/card/%s/showtoken", cardID)
	method := "POST"

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
