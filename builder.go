package ocmf_go

import (
	"crypto"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
)

type Builder struct {
	payload   PayloadSection
	signature Signature
}

func NewBuilder() *Builder {
	return &Builder{
		signature: Signature{
			Algorithm: string(SignatureAlgorithmECDSAsecp256r1SHA256),
			Encoding:  string(SignatureEncodingHex),
			MimeType:  string(SignatureMimeTypeDer),
		},
	}
}

func (b *Builder) WithFormatVersion(formatVersion string) *Builder {
	b.payload.FormatVersion = formatVersion
	return b
}

func (b *Builder) WithGatewayID(gatewayID string) *Builder {
	b.payload.GatewayID = gatewayID
	return b
}

func (b *Builder) WithGatewaySerial(gatewaySerial string) *Builder {
	b.payload.GatewaySerial = gatewaySerial
	return b
}

func (b *Builder) WithGatewayVersion(gatewayVersion string) *Builder {
	b.payload.GatewayVersion = gatewayVersion
	return b
}

func (b *Builder) WithPagination(pagination string) *Builder {
	b.payload.Pagination = pagination
	return b
}

// Sign payload
func (b *Builder) signPayload(privateKey crypto.PrivateKey) {
	var signedData string

	switch b.signature.Algorithm {
	case string(SignatureAlgorithmECDSAsecp192k1SHA256):
		// TODO
	case string(SignatureAlgorithmECDSAsecp256k1SHA256):
		// TODO
	case string(SignatureAlgorithmECDSAsecp384r1SHA256):
		// TODO
	case string(SignatureAlgorithmECDSAbrainpool256r11SHA256):
		// TODO
	case string(SignatureAlgorithmECDSAsecp256r1SHA256):
	// TODO
	default:

	}

	// Encode signed data
	switch b.signature.Encoding {
	case string(SignatureEncodingBase64):
		signedData = base64.StdEncoding.EncodeToString([]byte(signedData))
	default:
		signedData = fmt.Sprintf("%x", signedData)
	}

	b.signature.Data = signedData
}

// Sign payload
func (b *Builder) validatePayload() error {
	return nil
}

func (b *Builder) Build(privateKey crypto.PrivateKey) (*string, error) {
	err := b.validatePayload()
	if err != nil {
		return nil, err
	}

	// Build payload
	payload, err := json.Marshal(b.payload)
	if err != nil {
		return nil, err
	}

	// Sign payload
	b.signPayload(privateKey)

	// Build signature
	signature, err := json.Marshal(b.signature)
	if err != nil {
		return nil, err
	}

	return lo.ToPtr(fmt.Sprintf("OCMF|%v|%v", payload, signature)), nil
}
