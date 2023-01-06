package solid

import (
	"encoding/json"
	"fmt"

	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

func (service *SolidService) AddFamily(personID string, model models.FamilyDataPartial) (models.FamilyDataFull, error) {
	var response models.FamilyDataFull
	ctxError := "add-family"

	relativeURL := "/v1/family"
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
