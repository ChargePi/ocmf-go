package ocmf_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
