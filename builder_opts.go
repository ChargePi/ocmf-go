package ocmf_go

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

func WithSignature(signature Signature) BuilderOption {
	return func(b *Builder) {
		err := signature.Validate()
		if err != nil {
			return
		}

		b.signature = signature
	}
}
