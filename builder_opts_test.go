package ocmf_go

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/suite"
)

type builderOptsTestSuite struct {
	suite.Suite
}

func (s *builderOptsTestSuite) TestWithSignatureAlgorithm() {
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
			curve := elliptic.P256()
			privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
			s.Require().NoError(err)

			builder := NewBuilder(privateKey, WithSignatureAlgorithm(tt.algorithm))

			if tt.name == "Unknown algorithm" || tt.name == "Empty algorithm" {
				s.NotEqual(tt.algorithm, builder.signature.Algorithm)
			} else {
				s.Equal(tt.algorithm, builder.signature.Algorithm)
			}
		})
	}
}

func (s *builderOptsTestSuite) TestWithWithSignatureEncoding() {
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
			builder := NewBuilder(nil, WithSignatureEncoding(tt.encoding))

			if tt.encoding == SignatureEncodingBase64 || tt.encoding == SignatureEncodingHex {
				s.Equal(tt.encoding, builder.signature.Encoding)
			} else {
				s.NotEqual(tt.encoding, builder.signature.Encoding)
			}
		})
	}
}
func TestBuilderOpts(t *testing.T) {
	suite.Run(t, new(builderOptsTestSuite))
}
