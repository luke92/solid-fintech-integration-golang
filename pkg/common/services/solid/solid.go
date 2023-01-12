package solid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

type SolidService struct {
	BaseURL string
	APIKey  string
}

type sendRequestInput struct {
	PersonId    string
	Method      string
	RelativeURL string
	Payload     []byte
}

type sendRequestOutput struct {
	Status     int
	Data       []byte
	ErrorSolid *models.ErrorResponse
}

func NewSolidService(baseURL string, apiKey string) Service {
	return &SolidService{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

func (service *SolidService) sendRequest(input sendRequestInput) (sendRequestOutput, error) {
	output := sendRequestOutput{
		Status: http.StatusInternalServerError,
	}

	url := service.BaseURL + input.RelativeURL

	fmt.Printf("Request: %s : %s \n", input.Method, url)

	request, err := http.NewRequest(input.Method, url, bytes.NewBuffer(input.Payload))
	if err != nil {
		fmt.Println("Error preparing request")
		return output, err
	}

	request.Header.Add("sd-api-key", service.APIKey)
	if input.PersonId != "" {
		request.Header.Add("sd-person-id", input.PersonId)
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request")
		return output, err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	output.Data = body
	output.Status = response.StatusCode

	if response.StatusCode >= 400 {
		var errorResponse models.ErrorResponse
		err = json.Unmarshal(output.Data, &errorResponse)
		if err != nil {
			fmt.Println("Error parsing error from Solid")
			return output, err
		}

		fmt.Println("Error handled by Solid")
		output.ErrorSolid = &errorResponse
	}

	return output, nil
}
