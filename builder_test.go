package ocmf_go

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type builderTestSuite struct {
	suite.Suite
}

func (s *builderTestSuite) SetupTest() {
}

func (s *builderTestSuite) TearDownSuite() {
}

func (s *builderTestSuite) TestNewBuilder() {
	tests := []struct {
		name string
		opts []BuilderOption
	}{
		{},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {

		})
	}
}

func (s *builderTestSuite) TestBuilder_Valid() {
	privateKey := "" // todo
	builder := NewBuilder(privateKey).
		WithPagination("1").
		WithMeterSerial("exampleSerial123").
		WithIdentificationStatus(true).
		WithIdentificationType(string(RfidNone)).
		AddReading(Reading{
			Time:         "2018-07-24T13:22:04,000+0200",
			ReadingValue: 123,
			ReadingUnit:  string(UnitskWh),
			Status:       string(MeterOk),
		})

	s.Equal("1", builder.payload.Pagination)
	s.Equal("exampleSerial123", builder.payload.MeterSerial)
	s.Equal(true, builder.payload.IdentificationStatus)
	s.Equal(string(RfidNone), builder.payload.IdentificationType)
	s.Len(builder.payload.Readings, 1)
	s.Equal("2018-07-24T13:22:04,000+0200", builder.payload.Readings[0].Time)
	s.Equal(123, builder.payload.Readings[0].ReadingValue)
	s.Equal(string(UnitskWh), builder.payload.Readings[0].ReadingUnit)
	s.Equal(string(MeterOk), builder.payload.Readings[0].Status)

	payload, err := builder.Build()
	s.NoError(err)
	s.NotNil(payload)
}

func (s *builderTestSuite) TestBuilder_MissingAttributes() {
	builder := NewBuilder("privateKey").
		// WithPagination("1").
		WithMeterSerial("exampleSerial123").
		WithIdentificationStatus(true).
		WithIdentificationType(string(RfidNone)).
		AddReading(Reading{
			Time:         "2021-01-01T00:00:00Z",
			ReadingValue: 123,
			ReadingUnit:  string(UnitskWh),
			Status:       string(MeterOk),
		})

	payload, err := builder.Build()
	s.Error(err)
	s.Nil(payload)
}

func (s *builderTestSuite) TestBuilder_CantSign() {
	builder := NewBuilder("privateKey").
		WithPagination("1").
		WithMeterSerial("exampleSerial123").
		WithIdentificationStatus(true).
		WithIdentificationType(string(RfidNone)).
		AddReading(Reading{
			Time:         "2021-01-01T00:00:00Z",
			ReadingValue: 123,
			ReadingUnit:  string(UnitskWh),
			Status:       string(MeterOk),
		})

	builder.privateKey = ""

	payload, err := builder.Build()
	s.Error(err)
	s.Nil(payload)
}

func (s *builderTestSuite) TestWithSignatureAlgorithm() {
	tests := []struct {
		name      string
		algorithm SignatureAlgorithm
		want      bool
	}{
		{
			name:      "ECDSA-secp192k1-SHA256",
			algorithm: SignatureAlgorithmECDSAsecp192k1SHA256,
			want:      true,
		},
		{
			name:      "ECDSA-secp256k1-SHA256",
			algorithm: SignatureAlgorithmECDSAsecp256k1SHA256,
			want:      true,
		},
		{
			name:      "ECDSA-secp384r1-SHA256",
			algorithm: SignatureAlgorithmECDSAsecp384r1SHA256,
			want:      true,
		},
		{
			name:      "ECDSA-brainpool256r1-SHA256",
			algorithm: SignatureAlgorithmECDSAbrainpool256r11SHA256,
			want:      true,
		},
		{
			name:      "ECDSA-secp256r1-SHA256",
			algorithm: SignatureAlgorithmECDSAsecp256r1SHA256,
			want:      true,
		},
		{
			name:      "ECDSA-secp192r1-SHA256",
			algorithm: SignatureAlgorithmECDSAsecp192r1SHA256,
			want:      true,
		},
		{
			name:      "Unknown algorithm",
			algorithm: SignatureAlgorithm("ABCD"),
			want:      false,
		},
		{
			name:      "Empty algorithm",
			algorithm: SignatureAlgorithm(""),
			want:      false,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			builder := NewBuilder("privateKey", WithSignatureAlgorithm(tt.algorithm))

			if tt.name == "Unknown algorithm" || tt.name == "Empty algorithm" {
				s.NotEqual(tt.algorithm, builder.signature.Algorithm)
			} else {
				s.Equal(tt.algorithm, builder.signature.Algorithm)
			}
		})
	}
}

func (s *builderTestSuite) TestWithWithSignatureEncoding() {
	tests := []struct {
		name     string
		encoding SignatureEncoding
	}{
		{
			name:     "Base64 encoding",
			encoding: SignatureEncodingBase64,
		},
		{
			name:     "Hex encoding",
			encoding: SignatureEncodingHex,
		},
		{
			name:     "Empty encoding",
			encoding: SignatureEncoding(""),
		},
		{
			name:     "Unknown encoding",
			encoding: SignatureEncoding("ABDD"),
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			builder := NewBuilder("privateKey", WithSignatureEncoding(tt.encoding))

			if tt.encoding == SignatureEncodingBase64 || tt.encoding == SignatureEncodingHex {
				s.Equal(tt.encoding, builder.signature.Encoding)
			} else {
				s.NotEqual(tt.encoding, builder.signature.Encoding)
			}
		})
	}
}

func TestBuilder(t *testing.T) {
	suite.Run(t, new(builderTestSuite))
}
