package ocmf_go

import "crypto/ecdsa"

type ParserOpts struct {
	withAutomaticValidation            bool
	withAutomaticSignatureVerification bool
	publicKey                          *ecdsa.PublicKey
}

type Opt func(*ParserOpts)

func WithAutomaticValidation() Opt {
	return func(p *ParserOpts) {
		p.withAutomaticValidation = true
	}
}

func WithAutomaticSignatureVerification(publicKey *ecdsa.PublicKey) Opt {
	return func(p *ParserOpts) {
		p.withAutomaticSignatureVerification = true
		p.publicKey = publicKey
	}
}

func defaultOpts() ParserOpts {
	return ParserOpts{
		withAutomaticValidation: false,
	}
}
