package solid

import (
	"encoding/json"
	"fmt"

	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

func (service *SolidService) AddPerson(person models.PersonDataPartial) (models.PersonDataFull, error) {
	relativeURL := "/v1/person"
	method := "POST"
	payload, err := json.Marshal(person)
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
