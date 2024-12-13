package ocmf_go

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/suite"
)

type parserOptsTestSuite struct {
	suite.Suite
}

func (s *parserOptsTestSuite) TestParserOptions() {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	s.Require().NoError(err)

	tests := []struct {
		name            string
		opts            []Opt
		expectedOptions ParserOpts
	}{
		{
			name: "Default options",
			opts: []Opt{},
			expectedOptions: ParserOpts{
				withAutomaticValidation:            false,
				withAutomaticSignatureVerification: false,
				publicKey:                          nil,
			},
		},
		{
			name: "With automatic validation",
			opts: []Opt{
				WithAutomaticValidation(),
			},
			expectedOptions: ParserOpts{
				withAutomaticValidation:            true,
				withAutomaticSignatureVerification: false,
				publicKey:                          nil,
			},
		},
		{
			name: "With automatic signature verification but public key is empty",
			opts: []Opt{
				WithAutomaticSignatureVerification(nil),
			},
			expectedOptions: ParserOpts{
				withAutomaticValidation:            false,
				withAutomaticSignatureVerification: false,
				publicKey:                          nil,
			},
		},
		{
			name: "With automatic signature verification",
			opts: []Opt{
				WithAutomaticSignatureVerification(&privateKey.PublicKey),
			},
			expectedOptions: ParserOpts{
				withAutomaticValidation:            false,
				withAutomaticSignatureVerification: true,
				publicKey:                          &privateKey.PublicKey,
			},
		},
		{
			name: "With automatic validation and signature verification",
			opts: []Opt{
				WithAutomaticValidation(),
				WithAutomaticSignatureVerification(&privateKey.PublicKey),
			},
			expectedOptions: ParserOpts{
				withAutomaticValidation:            true,
				withAutomaticSignatureVerification: true,
				publicKey:                          &privateKey.PublicKey,
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			parser := NewParser(tt.opts...)
			s.Equal(tt.expectedOptions, parser.opts)
		})
	}
}

func (s *parserOptsTestSuite) TestParserDefaultOptions() {
	opts := defaultOpts()
	expectedDefaults := ParserOpts{
		withAutomaticValidation:            false,
		withAutomaticSignatureVerification: false,
		publicKey:                          nil,
	}
	s.Equal(expectedDefaults, opts)
}

func TestParserOpts(t *testing.T) {
	suite.Run(t, new(parserOptsTestSuite))
}
