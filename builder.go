package ocmf_go

import (
	"crypto"
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
)

type BuilderOption func(*Builder)

func WithSignatureAlgorithm(algorithm SignatureAlgorithm) BuilderOption {
	return func(b *Builder) {
		if isValidSignatureAlgorithm(algorithm) {
			b.signature.Algorithm = algorithm
		}
	}
}

func WithSignatureEncoding(encoding SignatureEncoding) BuilderOption {
	return func(b *Builder) {
		if isValidSignatureEncoding(encoding) {
			b.signature.Encoding = encoding
		}
	}
}

type Builder struct {
	payload    PayloadSection
	signature  Signature
	privateKey crypto.PrivateKey
}

func NewBuilder(privateKey crypto.PrivateKey, opts ...BuilderOption) *Builder {
	builder := &Builder{
		payload: PayloadSection{
			FormatVersion: OcmfVersion,
		},
		// Set default signature parameters
		signature:  *NewDefaultSignature(),
		privateKey: privateKey,
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

func (b *Builder) AddIdentificationFlag(flag string) *Builder {
	b.payload.IdentificationFlags = append(b.payload.IdentificationFlags, flag)
	return b
}

func (b *Builder) WithMeterSerial(serial string) *Builder {
	b.payload.MeterSerial = serial
	return b
}

func (b *Builder) WithIdentificationStatus(status bool) *Builder {
	b.payload.IdentificationStatus = status
	return b
}

func (b *Builder) WithIdentificationLevel(level string) *Builder {
	b.payload.IdentificationLevel = level
	return b
}

func (b *Builder) WithIdentificationType(idType string) *Builder {
	b.payload.IdentificationType = idType
	return b
}

func (b *Builder) WithIdentificationData(data string) *Builder {
	b.payload.IdentificationData = data
	return b
}

func (b *Builder) WithTariffText(text string) *Builder {
	b.payload.TariffText = text
	return b
}

func (b *Builder) WithChargeControllerVersion(version string) *Builder {
	b.payload.ChargeControllerVersion = version
	return b
}

func (b *Builder) WithChargePointIdentificationType(serial string) *Builder {
	b.payload.ChargePointIdentificationType = serial
	return b
}

func (b *Builder) WithChargePointIdentification(serial string) *Builder {
	b.payload.ChargePointIdentification = serial
	return b
}

func (b *Builder) AddLossCompensation(lossCompensation LossCompensation) *Builder {
	b.payload.LossCompensation = lossCompensation
	return b
}

func (b *Builder) ClearPayloadSection() *Builder {
	b.payload = PayloadSection{
		FormatVersion: OcmfVersion,
	}
	return b
}

func (b *Builder) Build() (*string, error) {
	// Validate payload
	err := b.payload.Validate()
	if err != nil {
		return nil, err
	}

	// Sign payload with private key
	err = b.signature.Sign(b.privateKey)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(b.payload)
	if err != nil {
		return nil, err
	}

	signature, err := json.Marshal(b.signature)
	if err != nil {
		return nil, err
	}

	return lo.ToPtr(fmt.Sprintf("OCMF|%v|%v", payload, signature)), nil
}
