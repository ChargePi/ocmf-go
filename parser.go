package ocmf_go

import (
	"encoding/json"
	"errors"
	"strings"
)

type Opt func(*Parser)

func WithAutomaticValidation() Opt {
	return func(p *Parser) {
		p.withAutomaticValidation = true
	}
}

var (
	ErrInvalidFormat = errors.New("invalid OCMF message format")
)

type Parser struct {
	payload                 *PayloadSection
	signature               *Signature
	withAutomaticValidation bool
	err                     error
}

func NewParser(opts ...Opt) *Parser {
	parser := &Parser{}

	// Apply opts
	for _, opt := range opts {
		opt(parser)
	}

	return parser
}

// Returns a new Parser instance with the payload and signature fields set
func (p *Parser) ParseOcmfMessageFromString(data string) *Parser {
	payloadSection, signature, err := parseOcmfMessageFromString(data)
	if err != nil {
		return &Parser{err: err}
	}

	return &Parser{
		payload:   payloadSection,
		signature: signature,
	}
}

func (p *Parser) GetPayload() (*PayloadSection, error) {
	if p.err != nil {
		return nil, p.err
	}

	// Validate the payload if automatic validation is enabled
	if p.withAutomaticValidation {
		if err := p.payload.Validate(); err != nil {
			return nil, err
		}
	}

	return p.payload, nil
}

func (p *Parser) GetSignature() (*Signature, error) {
	if p.err != nil {
		return nil, p.err
	}

	// Validate the signature if automatic validation is enabled
	if p.withAutomaticValidation {
		if err := p.signature.Validate(); err != nil {
			return nil, err
		}
	}

	return p.signature, nil
}

func parseOcmfMessageFromString(data string) (*PayloadSection, *Signature, error) {
	if !strings.HasPrefix(data, "OCMF|") {
		return nil, nil, ErrInvalidFormat
	}

	data, _ = strings.CutPrefix(data, "OCMF|")
	splitData := strings.Split(data, "|")

	if len(splitData) != 2 {
		return nil, nil, ErrInvalidFormat
	}

	payloadSection := &PayloadSection{}
	err := json.Unmarshal([]byte(splitData[0]), payloadSection)
	if err != nil {
		return nil, nil, err
	}

	signature := &Signature{}
	err = json.Unmarshal([]byte(splitData[1]), signature)
	if err != nil {
		return nil, nil, err
	}

	return payloadSection, signature, nil
}
