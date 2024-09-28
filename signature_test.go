package ocmf_go

import (
	"crypto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestIsValidSignatureEncoding(t *testing.T) {
	tests := []struct {
		name     string
		encoding SignatureEncoding
		want     bool
	}{
		{
			name:     "Base64 encoding",
			encoding: SignatureEncodingBase64,
			want:     true,
		},
		{
			name:     "Hex encoding",
			encoding: SignatureEncodingHex,
			want:     true,
		},
		{
			name:     "Empty encoding",
			encoding: SignatureEncoding(""),
			want:     false,
		},
		{
			name:     "Unknown encoding",
			encoding: SignatureEncoding("ABDD"),
			want:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isValidSignatureEncoding(test.encoding)
			assert.Equal(t, test.want, res)
		})
	}
}

func TestIsValidSignatureAlgorithm(t *testing.T) {
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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isValidSignatureAlgorithm(test.algorithm)
			assert.Equal(t, test.want, res)
		})
	}
}

type signatureTestSuite struct {
	suite.Suite
}

func (s *signatureTestSuite) SetupTest() {}

func (s *signatureTestSuite) TestValidate() {
	tests := []struct {
		name      string
		signature Signature
		error     bool
	}{
		{
			name: "Valid signature",
			signature: Signature{
				Algorithm: SignatureAlgorithmECDSAsecp256r1SHA256,
				Encoding:  SignatureEncodingHex,
				MimeType:  SignatureMimeTypeDer,
				Data:      "data",
			},
		},
		{
			name: "Invalid encoding",
			signature: Signature{
				Algorithm: SignatureAlgorithmECDSAsecp256r1SHA256,
				Encoding:  "",
				MimeType:  SignatureMimeTypeDer,
				Data:      "data",
			},
			error: true,
		},
		{
			name: "Invalid algorithm",
			signature: Signature{
				Algorithm: "",
				Encoding:  SignatureEncodingHex,
				MimeType:  SignatureMimeTypeDer,
				Data:      "data",
			},
			error: true,
		},
		{
			name: "Empty data",
			signature: Signature{
				Algorithm: SignatureAlgorithmECDSAsecp256r1SHA256,
				Encoding:  SignatureEncodingHex,
				MimeType:  SignatureMimeTypeDer,
				Data:      "",
			},
			error: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			err := tt.signature.Validate()
			if tt.error {
				s.Error(err)
			} else {
				s.NoError(err)
			}
		})
	}
}

func (s *signatureTestSuite) TestSign() {
	tests := []struct {
		name       string
		signature  Signature
		privateKey crypto.PrivateKey
		publicKey  crypto.PublicKey
		error      bool
	}{
		{
			name: "Valid signature",
		},
		{
			name: "Invalid algorithm",
		},
		{
			name: "Invalid encoding",
		},
		{
			name: "Invalid private key",
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			err := tt.signature.Sign(tt.privateKey)
			if tt.error {
				s.Error(err)
			} else {
				s.NoError(err)
			}

			err = tt.signature.Verify(tt.publicKey)
			if tt.error {
				s.Error(err)
			} else {
				s.NoError(err)
			}
		})
	}
}

func TestSignature(t *testing.T) {
	suite.Run(t, new(signatureTestSuite))
}
