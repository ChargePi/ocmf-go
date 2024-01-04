package ocmf_go

type Signature struct {
	Algorithm string `json:"SA"`
	Encoding  string `json:"SE,omitempty"`
	MimeType  string `json:"SM,omitempty"`
	Data      string `json:"SD"`
}

type SignatureMimeType string

const (
	SignatureMimeTypeDer = SignatureMimeType("application/x-der")
)

type SignatureEncoding string

const (
	SignatureEncodingBase64 = SignatureEncoding("base64")
	SignatureEncodingHex    = SignatureEncoding("hex")
)

type SignatureAlgorithm string

const (
	SignatureAlgorithmECDSAsecp192k1SHA256       = SignatureAlgorithm("ECDSA-secp192k1-SHA256")
	SignatureAlgorithmECDSAsecp256k1SHA256       = SignatureAlgorithm("ECDSA-secp256k1-SHA256")
	SignatureAlgorithmECDSAsecp384r1SHA256       = SignatureAlgorithm("ECDSA-secp384r1-SHA256")
	SignatureAlgorithmECDSAbrainpool256r11SHA256 = SignatureAlgorithm("ECDSA-brainpool256r1-SHA256")
	SignatureAlgorithmECDSAsecp256r1SHA256       = SignatureAlgorithm("ECDSA-secp256r1-SHA256")
	SignatureAlgorithmECDSAsecp192r1SHA256       = SignatureAlgorithm("ECDSA-secp192r1-SHA256")
)
