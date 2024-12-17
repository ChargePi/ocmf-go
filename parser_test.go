package ocmf_go

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
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

func (s *parserTestSuite) TestGetPayload_valid() {
	parser := NewParser().ParseOcmfMessageFromString(examplePayload)

	payload, err := parser.GetPayload()
	s.NoError(err)
	s.NotNil(payload)

	s.Equal("EEM-350-D-MCB", payload.MeterModel)
	s.Equal("BQ27400330016", payload.MeterSerial)
}

func (s *parserTestSuite) TestGetPayload_unparsable() {
	malformedPayload := "OCMF|{}|{"
	parser := NewParser().ParseOcmfMessageFromString(malformedPayload)

	payload, err := parser.GetPayload()
	s.ErrorContains(err, "failed to unmarshal signature")
	s.Nil(payload)
}

func (s *parserTestSuite) TestGetSignature_valid() {
	// Generate private and public ECDSA keys
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	s.Require().NoError(err)

	builder := NewBuilder(privateKey).
		WithPagination("1").
		WithMeterSerial("exampleSerial123").
		WithIdentificationStatus(true).
		WithIdentificationType(string(RfidNone)).
		AddReading(Reading{
			Time:         "2018-07-24T13:22:04,000+0200 S",
			ReadingValue: 1.0,
			ReadingUnit:  string(UnitskWh),
			Status:       string(MeterOk),
		})

	message, err := builder.Build()
	s.Require().NoError(err)

	tests := []struct {
		name              string
		parserOpts        []Opt
		data              string
		expectedSignature *Signature
	}{
		{
			name:              "No validation",
			parserOpts:        []Opt{},
			data:              *message,
			expectedSignature: &builder.signature,
		},
		{
			name: "With automatic signature verification",
			parserOpts: []Opt{
				WithAutomaticSignatureVerification(&privateKey.PublicKey),
			},
			data:              *message,
			expectedSignature: &builder.signature,
		},
		{
			name: "With automatic payload validation",
			parserOpts: []Opt{
				WithAutomaticValidation(),
			},
			data:              *message,
			expectedSignature: &builder.signature,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			parser := NewParser(tt.parserOpts...).ParseOcmfMessageFromString(tt.data)

			signature, err := parser.GetSignature()
			s.NoError(err)
			s.NotNil(signature)
			s.Equal(*tt.expectedSignature, *signature)
		})
	}
}

func (s *parserTestSuite) TestGetSignature_invalid() {
	// Generate private and public ECDSA keys
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	s.Require().NoError(err)

	builder := NewBuilder(privateKey).
		WithPagination("1").
		WithMeterSerial("exampleSerial123").
		WithIdentificationStatus(true).
		WithIdentificationType(string(RfidNone)).
		AddReading(Reading{
			Time:         "2018-07-24T13:22:04,000+0200 S",
			ReadingValue: 1.0,
			ReadingUnit:  string(UnitskWh),
			Status:       string(MeterOk),
		})

	message, err := builder.Build()
	s.Require().NoError(err)

	privateKey2, err := ecdsa.GenerateKey(curve, rand.Reader)
	s.Require().NoError(err)

	tests := []struct {
		name       string
		parserOpts []Opt
		data       string
		error      string
	}{
		{
			name: "Signature validation failed",
			parserOpts: []Opt{
				WithAutomaticSignatureVerification(&privateKey2.PublicKey),
			},
			data:  *message,
			error: "verification failed",
		}, {
			name: "Nil public key",
			parserOpts: []Opt{
				WithAutomaticSignatureVerification(nil),
			},
			data:  *message,
			error: "unable to verify signature",
		},
		{
			name: "Payload empty",
			parserOpts: []Opt{
				WithAutomaticSignatureVerification(&privateKey.PublicKey),
			},
			data:  *message,
			error: "payload is empty",
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			parser := NewParser(tt.parserOpts...).ParseOcmfMessageFromString(tt.data)

			if tt.name == "Payload empty" {
				parser.payload = nil
			}

			_, err := parser.GetSignature()
			s.ErrorContains(err, tt.error)
		})
	}
}

func TestParser(t *testing.T) {
	suite.Run(t, new(parserTestSuite))
}
