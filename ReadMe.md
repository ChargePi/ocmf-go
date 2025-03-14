# OCMF-go

![coverage](https://raw.githubusercontent.com/ChargePi/ocmf-go/badges/.badges/main/coverage.svg)

OCMF-go is an implementation of the Open Charge Metering Format (OCMF) in Go. It provides a simple library for
generating and parsing OCMF messages. The provided message builder generates OCMF-compatible messages and signs the data
with the provided private key, desired algorithm and encoding, so you don't have to.

## Installation

```shell
go get github.com/ChargePi/ocmf-go
```

## Usage

```go
package main

import (
	"fmt"

	ocmf_go "github.com/ChargePi/ocmf-go"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
)

func main() {
	// Generate a new message builder
	builder := ocmf_go.NewBuilder()

	// Set the signature algorithm
	builder.SetSignatureAlgorithm(ocmf_go.SignatureAlgorithmECDSAsecp256r1Sha256)

	// Set the signature encoding
	builder.SetSignatureEncoding(ocmf_go.SignatureEncodingBase64)

	// ... set the desired fields
	message, err := builder.Build()
	if err != nil {
		fmt.Println(err)
	}

	// Create a MeterValue message with the generated message as value 
	meterValueExample := types.MeterValue{
		SampledValue: []types.SampledValue{
			{
				Value:  message,
				Format: types.ValueFormatSignedData,
			},
		},
	}

	// Send the message via OCPP 1.5/1.6.
}

```

## Contributing

Contributions are welcome! Please check out the [contributing guide](/docs/contributing/contributing.md) for more
information.

## License

OCMF-go is licensed under the [MIT License](LICENSE.txt).