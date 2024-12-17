package ocmf_go

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrInvalidFormat       = errors.New("invalid OCMF message format")
	ErrVerificationFailure = errors.New("verification failed")
	ErrPayloadEmpty        = errors.New("payload is empty")
)

type Parser struct {
	payload   *PayloadSection
	signature *Signature
	opts      ParserOpts
	err       error
}

func NewParser(opts ...Opt) *Parser {
	defaults := defaultOpts()
	// Apply opts
	for _, opt := range opts {
		opt(&defaults)
	}

	return &Parser{
		opts: defaults,
	}
}

// ParseOcmfMessageFromString Returns a new Parser instance with the payload and signature fields set
func (p *Parser) ParseOcmfMessageFromString(data string) *Parser {
	payloadSection, signature, err := parseOcmfMessageFromString(data)
	if err != nil {
		return &Parser{err: err, opts: p.opts}
	}

	return &Parser{
		payload:   payloadSection,
		signature: signature,
		opts:      p.opts,
	}
}

func (p *Parser) GetPayload() (*PayloadSection, error) {
	if p.err != nil {
		return nil, p.err
	}

	// Validate the payload if automatic validation is enabled
	if p.opts.withAutomaticValidation {
		if err := p.payload.Validate(); err != nil {
			return nil, errors.Wrap(err, "payload validation failed")
		}
	}

	return p.payload, nil
}

func (p *Parser) GetSignature() (*Signature, error) {
	if p.err != nil {
		return nil, p.err
	}

	// Validate the signature if automatic validation is enabled
	if p.opts.withAutomaticValidation {
		if err := p.signature.Validate(); err != nil {
			return nil, errors.Wrap(err, "signature validation failed")
		}
	}

	if p.opts.withAutomaticSignatureVerification {
		if p.payload == nil {
			return nil, ErrPayloadEmpty
		}

		valid, err := p.signature.Verify(*p.payload, p.opts.publicKey)
		if err != nil {
			return nil, errors.Wrap(err, "unable to verify signature")
		}

		// Even if the signature is valid, we still return an error if the verification failed
		if !valid {
			return p.signature, ErrVerificationFailure
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

	payloadSection := PayloadSection{}
	err := json.Unmarshal([]byte(splitData[0]), &payloadSection)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshal payload")
	}

	signature := Signature{}
	err = json.Unmarshal([]byte(splitData[1]), &signature)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshal signature")
	}

	return &payloadSection, &signature, nil
}
