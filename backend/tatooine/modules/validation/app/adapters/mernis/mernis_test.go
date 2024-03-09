package mernis

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMernis_parseSOAPResponse_Valid(t *testing.T) {

	soapResponse := `
<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope">
  <soap:Body>
	<TCKimlikNoDogrulaResponse xmlns="http://tckimlik.nvi.gov.tr/WS">
	  <TCKimlikNoDogrulaResult>true</TCKimlikNoDogrulaResult>
	</TCKimlikNoDogrulaResponse>
  </soap:Body>
</soap:Envelope>
`
	reader := strings.NewReader(soapResponse)

	result, err := parseSOAPResponse(reader)

	assert.NoError(t, err)
	assert.True(t, result.Body.TCKimlikNoDogrulaResponse.Result)

}

func TestMernis_parseSOAPResponse_Invalid(t *testing.T) {
	soapResponse := `
<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope">
  <soap:Body>
	<TCKimlikNoDogrulaResponse xmlns="http://tckimlik.nvi.gov.tr/WS">
	  <TCKimlikNoDogrulaResult>true</TCKimlikNoDogrulaResult>
	</TCKimlikNoDogrulaResponse>
  </soap:Body>
</soap:Envelope
`
	reader := strings.NewReader(soapResponse)

	result, err := parseSOAPResponse(reader)

	assert.Error(t, err)
	assert.False(t, result.Body.TCKimlikNoDogrulaResponse.Result)
}
