package veriff

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/domain/model"
)

type veriffAdapter struct {
	creds model.VeriffAPICred
}

func NewVeriffAdapter() *veriffAdapter {
	apiKey := os.Getenv("VERIFF_API_KEY")
	secretKey := os.Getenv("VERIFF_API_KEY")
	baseUrl := os.Getenv("VERIFF_API_KEY")

	return &veriffAdapter{
		creds: model.VeriffAPICred{
			ApiKey:    apiKey,
			SecretKey: secretKey,
			BaseUrl:   baseUrl,
		},
	}
}

func (v *veriffAdapter) GenerateSession() (string, error) {
	veriffResponse, err := v.makeRequestToVeriff()

	if veriffResponse.Status != "created" {
		return "", err
	}

	return veriffResponse.Url, nil
}

func (v *veriffAdapter) makeRequestToVeriff() (*model.VeriffResponse, error) {
	request := model.VeriffRequest{
		Callback:   "",
		VendorData: "",
	}

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed marshalling request: %v", err)
	}

	mac := hmac.New(sha256.New, []byte(v.creds.SecretKey))
	mac.Write(jsonRequest)
	signedPayload := mac.Sum(nil)

	agent := fiber.AcquireAgent()
	agent.Request().Header.SetMethod("POST")
	agent.Request().Header.Add("X-AUTH-CLIENT", v.creds.ApiKey)
	agent.Request().Header.SetContentType("application/json")
	agent.Request().SetRequestURI(v.creds.BaseUrl)
	agent.Request().Header.Add("X-HMAC-SIGNATURE", string(signedPayload))
	agent.Body(jsonRequest)

	err = agent.Parse()
	if err != nil {
		return nil, fmt.Errorf("unable to parse agent")
	}

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to send request")
	}
	var veriffResponse model.VeriffResponse
	err = json.Unmarshal(body, &veriffResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response")
	}

	return &veriffResponse, nil
}
