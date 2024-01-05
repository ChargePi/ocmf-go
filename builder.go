package ocmf_go

import (
	"crypto"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
)

type BuilderOption func(*Builder)

func WithSignatureAlgorithm(algorithm SignatureAlgorithm) BuilderOption {
	return func(b *Builder) {
		b.signature.Algorithm = algorithm
	}
}

func WithSignatureEncoding(encoding SignatureEncoding) BuilderOption {
	return func(b *Builder) {
		b.signature.Encoding = encoding
	}
}

type Builder struct {
	payload   PayloadSection
	signature Signature
}

func NewBuilder(opts ...BuilderOption) *Builder {
	builder := &Builder{
		payload: PayloadSection{
			FormatVersion: "0.4",
		},
		// Set default signature parameters
		signature: Signature{
			Algorithm: SignatureAlgorithmECDSAsecp256r1SHA256,
			Encoding:  SignatureEncodingHex,
			MimeType:  SignatureMimeTypeDer,
		},
	}

	// Apply builder options
	for _, option := range opts {
		option(builder)
	}

	return builder
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

func (b *Builder) AddReading(reading Reading) *Builder {
	b.payload.Readings = append(b.payload.Readings, reading)
	return b
}

func (b *Builder) AddFlag(flag string) *Builder {
	b.payload.IdentificationFlags = append(b.payload.IdentificationFlags, flag)
	return b
}

func (b *Builder) AddLossCompensation(lossCompensation LossCompensation) *Builder {
	b.payload.LossCompensation = lossCompensation
	return b
}

func (b *Builder) ClearPayloadSection() *Builder {
	b.payload = PayloadSection{
		FormatVersion: "0.4",
	}
	return b
}

// Sign payload
func (b *Builder) signPayload(privateKey crypto.PrivateKey) {
	var signedData string

	switch b.signature.Algorithm {
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

	}

	// Encode signed data
	switch b.signature.Encoding {
	case SignatureEncodingBase64:
		signedData = base64.StdEncoding.EncodeToString([]byte(signedData))
	default:
		signedData = hex.EncodeToString([]byte(signedData))
	}

	b.signature.Data = signedData
}

// Sign payload
func (b *Builder) validatePayload() error {
	return b.payload.Validate()
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
