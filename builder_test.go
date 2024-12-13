package ocmf_go

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
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
			ReadingValue: 123,
			ReadingUnit:  string(UnitskWh),
			Status:       string(MeterOk),
		})

	s.Equal("1", builder.payload.Pagination)
	s.Equal("exampleSerial123", builder.payload.MeterSerial)
	s.Equal(true, builder.payload.IdentificationStatus)
	s.Equal(string(RfidNone), builder.payload.IdentificationType)
	s.Len(builder.payload.Readings, 1)
	s.Equal("2018-07-24T13:22:04,000+0200 S", builder.payload.Readings[0].Time)
	s.Equal(float64(123), builder.payload.Readings[0].ReadingValue)
	s.Equal(string(UnitskWh), builder.payload.Readings[0].ReadingUnit)
	s.Equal(string(MeterOk), builder.payload.Readings[0].Status)

	payload, err := builder.Build()
	s.NoError(err)
	s.NotNil(payload)
}

func (s *builderTestSuite) TestBuilder_MissingAttributes() {
	builder := NewBuilder(nil).
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
	builder := NewBuilder(nil).
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

	payload, err := builder.Build()
	s.Error(err)
	s.Nil(payload)
}

func TestBuilder(t *testing.T) {
	suite.Run(t, new(builderTestSuite))
}
