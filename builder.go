package ocmf_go

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type Builder struct {
	payload    PayloadSection
	signature  Signature
	privateKey *ecdsa.PrivateKey
}

func NewBuilder(privateKey *ecdsa.PrivateKey, opts ...BuilderOption) *Builder {
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

func (b *Builder) WithMeterModel(model string) *Builder {
	b.payload.MeterModel = model
	return b
}

func (b *Builder) WithMeterVendor(vendor string) *Builder {
	b.payload.MeterVendor = vendor
	return b
}

func (b *Builder) WithMeterFirmware(firmware string) *Builder {
	b.payload.MeterFirmware = firmware
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

func (b *Builder) ClearPayloadSection() {
	b.payload = PayloadSection{
		FormatVersion: OcmfVersion,
	}
}

func (b *Builder) Build() (*string, error) {
	// Validate payload
	err := b.payload.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "payload validation failed")
	}

	// Sign payload with private key
	err = b.signature.Sign(b.payload, b.privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	payload, err := json.Marshal(b.payload)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}

	signature, err := json.Marshal(b.signature)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal signature")
	}

	return lo.ToPtr(fmt.Sprintf("OCMF|%s|%s", payload, signature)), nil
}
