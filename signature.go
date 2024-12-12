package ocmf_go

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
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

func (s *Signature) Sign(payload PayloadSection, privateKey *ecdsa.PrivateKey) error {
	if privateKey == nil {
		return errors.New("private key is required")
	}

	// Marshal payload
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal payload")
	}

	switch s.Algorithm {
	case SignatureAlgorithmECDSAsecp192k1SHA256:
	case SignatureAlgorithmECDSAsecp256k1SHA256:
	case SignatureAlgorithmECDSAsecp384r1SHA256:
	case SignatureAlgorithmECDSAbrainpool256r11SHA256:
	case SignatureAlgorithmECDSAsecp256r1SHA256:
	default:
		return fmt.Errorf("unsupported signature algorithm: %s", s.Algorithm)
	}

	// Hash data
	messageHash := sha256.Sum256(payloadBytes)

	// Sign data
	sign, err := ecdsa.SignASN1(rand.Reader, privateKey, messageHash[:])
	if err != nil {
		return errors.Wrap(err, "failed to sign data")
	}

	var signedData string

	// Encode signed data
	switch s.Encoding {
	case SignatureEncodingBase64:
		signedData = base64.StdEncoding.EncodeToString(sign)
	case SignatureEncodingHex:
		signedData = hex.EncodeToString(sign)
	default:
		return fmt.Errorf("unsupported signature encoding: %s", s.Encoding)
	}

	s.Data = signedData
	return nil
}

func (s *Signature) Verify(payload PayloadSection, publicKey *ecdsa.PublicKey) (bool, error) {
	var decoded []byte

	if publicKey == nil {
		return false, errors.New("public key is required")
	}

	// Decode the signature
	switch s.Encoding {
	case SignatureEncodingBase64:
		decodedString, err := base64.StdEncoding.DecodeString(s.Data)
		if err != nil {
			return false, errors.Wrap(err, "failed to decode base64 data")
		}

		decoded = decodedString
	case SignatureEncodingHex:
		decodedString, err := hex.DecodeString(s.Data)
		if err != nil {
			return false, errors.Wrap(err, "failed to decode hex data")
		}

		decoded = decodedString
	default:
		return false, fmt.Errorf("unsupported signature encoding: %s", s.Encoding)
	}

	switch s.Algorithm {
	case SignatureAlgorithmECDSAsecp192k1SHA256:
	case SignatureAlgorithmECDSAsecp256k1SHA256:
	case SignatureAlgorithmECDSAsecp384r1SHA256:
	case SignatureAlgorithmECDSAbrainpool256r11SHA256:
	case SignatureAlgorithmECDSAsecp256r1SHA256:
	default:
		return false, fmt.Errorf("unsupported signature algorithm: %s", s.Algorithm)
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return false, errors.Wrap(err, "failed to marshal payload")
	}

	// Hash the payload to compare with the signature
	messageHash := sha256.Sum256(payloadBytes)

	// Verify signature
	return ecdsa.VerifyASN1(publicKey, messageHash[:], decoded), nil
}
