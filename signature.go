package ocmf_go

import (
	"crypto"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

type SignatureMimeType string

const (
	SignatureMimeTypeDer = SignatureMimeType("application/x-der")
)

type SignatureEncoding string

const (
	SignatureEncodingBase64 = SignatureEncoding("base64")
	SignatureEncodingHex    = SignatureEncoding("hex")
)

func isValidSignatureEncoding(encoding SignatureEncoding) bool {
	switch encoding {
	case SignatureEncodingBase64, SignatureEncodingHex:
		return true
	default:
		return false
	}
}

type SignatureAlgorithm string

const (
	SignatureAlgorithmECDSAsecp192k1SHA256       = SignatureAlgorithm("ECDSA-secp192k1-SHA256")
	SignatureAlgorithmECDSAsecp256k1SHA256       = SignatureAlgorithm("ECDSA-secp256k1-SHA256")
	SignatureAlgorithmECDSAsecp384r1SHA256       = SignatureAlgorithm("ECDSA-secp384r1-SHA256")
	SignatureAlgorithmECDSAbrainpool256r11SHA256 = SignatureAlgorithm("ECDSA-brainpool256r1-SHA256")
	SignatureAlgorithmECDSAsecp256r1SHA256       = SignatureAlgorithm("ECDSA-secp256r1-SHA256")
	SignatureAlgorithmECDSAsecp192r1SHA256       = SignatureAlgorithm("ECDSA-secp192r1-SHA256")
)

func isValidSignatureAlgorithm(algorithm SignatureAlgorithm) bool {
	switch algorithm {
	case SignatureAlgorithmECDSAsecp192k1SHA256,
		SignatureAlgorithmECDSAsecp256k1SHA256,
		SignatureAlgorithmECDSAsecp384r1SHA256,
		SignatureAlgorithmECDSAbrainpool256r11SHA256,
		SignatureAlgorithmECDSAsecp256r1SHA256,
		SignatureAlgorithmECDSAsecp192r1SHA256:
		return true
	default:
		return false
	}
}

type Signature struct {
	Algorithm SignatureAlgorithm `json:"SA" validate:"required,signatureAlgorithm"`
	Encoding  SignatureEncoding  `json:"SE,omitempty" validate:"required,signatureEncoding"`
	MimeType  SignatureMimeType  `json:"SM,omitempty" validate:"required,oneof=application/x-der"`
	Data      string             `json:"SD" validate:"required"`
}

func NewDefaultSignature() *Signature {
	return &Signature{
		Algorithm: SignatureAlgorithmECDSAsecp256r1SHA256,
		Encoding:  SignatureEncodingHex,
		MimeType:  SignatureMimeTypeDer,
	}
}

func (s *Signature) Validate() error {
	return signatureValidator.Struct(s)
}

func (s *Signature) Sign(privateKey crypto.PrivateKey) error {
	var signedData string

	switch s.Algorithm {
	case SignatureAlgorithmECDSAsecp192k1SHA256:
		// TODO
	case SignatureAlgorithmECDSAsecp256k1SHA256:
		// TODO
	case SignatureAlgorithmECDSAsecp384r1SHA256:
		// TODO
	case SignatureAlgorithmECDSAbrainpool256r11SHA256:
		// TODO
	case SignatureAlgorithmECDSAsecp256r1SHA256:
	// TODO
	default:
		return fmt.Errorf("unsupported signature algorithm: %s", s.Algorithm)
	}

	// Encode signed data
	switch s.Encoding {
	case SignatureEncodingBase64:
		signedData = base64.StdEncoding.EncodeToString([]byte(signedData))
	case SignatureEncodingHex:
		signedData = hex.EncodeToString([]byte(signedData))
	default:
		return fmt.Errorf("unsupported signature encoding: %s", s.Encoding)
	}

	s.Data = signedData
	return nil
}
