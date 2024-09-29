package ocmf_go

import (
	"encoding/json"
	"errors"
	"strings"
)

var (
	ErrInvalidFormat = errors.New("invalid OCMF message format")
)

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

	// Validate the message
	err = payloadSection.Validate()
	if err != nil {
		return nil, nil, err
	}

	err = signature.Validate()
	if err != nil {
		return nil, nil, err
	}

	return payloadSection, signature, nil
}
