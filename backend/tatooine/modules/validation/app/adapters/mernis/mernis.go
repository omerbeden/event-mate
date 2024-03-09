package mernis

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type mernisAdapter struct {
	nationalId string
	name       string
	lastName   string
	birthYear  int
}

func NewMernisAdapter(nationalId, name, lastname string, birthYear int) *mernisAdapter {
	return &mernisAdapter{
		nationalId: nationalId,
		name:       name,
		lastName:   lastname,
		birthYear:  birthYear,
	}
}

func parseSOAPResponse(reader io.Reader) (*envelope, error) {
	var envelope envelope
	if err := xml.NewDecoder(reader).Decode(&envelope); err != nil {
		return nil, err
	}
	return &envelope, nil
}

func (mernis mernisAdapter) prepareSoapRequest() (*http.Request, error) {
	soapURL := "https://tckimlik.nvi.gov.tr/Service/KPSPublic.asmx?WSDL"
	soapBody := fmt.Sprintf(
		`<?xml version="1.0" encoding="utf-8"?>
	<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
	  <soap12:Body>
		<TCKimlikNoDogrula xmlns="http://tckimlik.nvi.gov.tr/WS">
		  <TCKimlikNo>%s</TCKimlikNo>
		  <Ad>%s</Ad>
		  <Soyad>%s</Soyad>
		  <DogumYili>%d</DogumYili>
		</TCKimlikNoDogrula>
	  </soap12:Body>
	</soap12:Envelope>`, mernis.nationalId, mernis.name, mernis.lastName, mernis.birthYear)

	req, err := http.NewRequest("POST", soapURL, bytes.NewBufferString(soapBody))
	if err != nil {
		return nil, fmt.Errorf("mernis prepare request: %w", err)
	}
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	return req, nil
}

func (mernis mernisAdapter) Validate() (bool, error) {

	req, err := mernis.prepareSoapRequest()
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("mernis client do: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("mernis read body: %w", err)
	}

	envelope, err := parseSOAPResponse(bytes.NewReader(bodyBytes))
	if err != nil {
		return false, fmt.Errorf("mernis parsing soap response: %w", err)
	}

	return envelope.Body.TCKimlikNoDogrulaResponse.Result, nil

}

type envelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Body    body     `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
}

type body struct {
	TCKimlikNoDogrulaResponse mernisResponse `xml:"http://tckimlik.nvi.gov.tr/WS TCKimlikNoDogrulaResponse"`
}

type mernisResponse struct {
	Result bool `xml:"TCKimlikNoDogrulaResult"`
}
