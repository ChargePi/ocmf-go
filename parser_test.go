package ocmf_go

import "testing"

func TestParseOcmfFromString(t *testing.T) {
	tests := []struct {
		name              string
		expectedPayload   PayloadSection
		expectedSignature Signature
		expectedError     error
	}{
		{
			name: "valid OCMF message",
		},
		{
			name: "Missing OCMF prefix",
		},
		{
			name: "Invalid OCMF format",
		},
		{
			name: "Invalid JSON in payload",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
