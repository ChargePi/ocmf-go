package ocmf_go

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

var examplePayload = "OCMF|{\n \"FV\": \"1.0\",\n \"GI\": \"ABL SBC-301\",\n \"GS\": \"808829900001\",\n \"GV\": \"1.4p3\",\n \"PG\": \"T12345\",\n \"MV\": \"Phoenix Contact\",\n \"MM\": \"EEM-350-D-MCB\",\n \"MS\": \"BQ27400330016\",\n \"MF\": \"1.0\",\n \"IS\": true,\n \"IL\": \"VERIFIED\",\n \"IF\": [\n \"RFID_PLAIN\",\n \"OCPP_RS_TLS\"\n ],\n \"IT\": \"ISO14443\",\n \"ID\": \"1F2D3A4F5506C7\",\n \"TT\": \"Tarif 1\",\n \"LC\": {\n \"LN\": \"cable_name\",\n \"LI\": 1,\n \"LR\": 2,\n \"LU\": \"mOhm\"\n },\n \"RD\": [\n {\n \"TM\": \"2018-07-24T13:22:04,000+0200 S\",\n \"TX\": \"B\",\n \"RV\": 2935.6,\n \"RI\": \"1-b:1.8.0\",\n \"RU\": \"kWh\",\n \"RT\": \"DC\",\n \"EF\": \"\",\n \"ST\": \"G\"\n },\n {\n \"TM\": \"2018-07-24T13:26:04,000+0200 S\",\n \"TX\": \"E\",\n \"RV\": 2965.1,\n \"CL\": 0.5,\n \"RI\": \"1-b:1.8.0\",\n \"RU\": \"kWh\",\n \"RT\": \"DC\",\n \"EF\": \"\",\n \"ST\": \"G\"\n }\n ]\n}|{\n \"SD\": \"887FABF407AC82782EEFFF2220C2F856AEB0BC22364BBCC6B55761911ED651D1A922BADA88818C9671AFEE7094D7F536\"\n}"

type parserTestSuite struct {
	suite.Suite
}

func (s *parserTestSuite) TestParseOcmfMessageFromString_valid() {
	payload, signature, err := parseOcmfMessageFromString(examplePayload)
	s.NoError(err)
	s.NotNil(payload)
	s.NotNil(signature)
}

func (s *parserTestSuite) TestParseOcmfMessageFromString_invalid_format() {
	payload, signature, err := parseOcmfMessageFromString("OCMF|{}|{data}")
	s.ErrorContains(err, "failed to unmarshal signature")
	s.Nil(payload)
	s.Nil(signature)

	payloadWithoutOCMF := strings.Replace(examplePayload, "OCMF|", "", 1)
	payload, signature, err = parseOcmfMessageFromString(payloadWithoutOCMF)
	s.ErrorIs(err, ErrInvalidFormat)
	s.Nil(payload)
	s.Nil(signature)

	malformedJsonPayload := strings.Replace(examplePayload, "}", "", 1)
	payload, signature, err = parseOcmfMessageFromString(malformedJsonPayload)
	s.ErrorContains(err, "failed to unmarshal payload")
	s.Nil(payload)
	s.Nil(signature)
}

func (s *parserTestSuite) TestGetPayload() {
	tests := []struct {
		name string
	}{}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {

		})
	}
}

func (s *parserTestSuite) TestGetSignature_valid() {
	tests := []struct {
		name              string
		parserOpts        []Opt
		data              string
		expectedSignature *Signature
	}{
		{
			name: "With automatic signature verification",
		},
		{
			name: "With automatic signature validation",
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {

		})
	}
}

func (s *parserTestSuite) TestGetSignature_invalid() {
	tests := []struct {
		name              string
		parserOpts        []Opt
		data              string
		expectedSignature *Signature
		error             string
	}{
		{},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			parser := NewParser(tt.parserOpts...)

			signature, err := parser.GetSignature()
			if tt.error != "" {
				s.ErrorContains(err, tt.error)
			} else {
				s.NoError(err)
				s.NotNil(signature)
				s.Equal(*tt.expectedSignature, *signature)
			}
		})
	}
}

func TestParser(t *testing.T) {
	suite.Run(t, new(parserTestSuite))
}
