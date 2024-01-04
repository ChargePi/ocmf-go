package ocmf_go

type PayloadSection struct {
	// General information
	FormatVersion  string `json:"FV,omitempty"`
	GatewayID      string `json:"GI,omitempty"`
	GatewaySerial  string `json:"GS,omitempty"`
	GatewayVersion string `json:"GV,omitempty"`
	// Pagination
	Pagination string `json:"PG"`
	// Meter identification
	MeterVendor   string `json:"MV,omitempty"`
	MeterModel    string `json:"MM,omitempty"`
	MeterSerial   string `json:"MS"`
	MeterFirmware string `json:"MF,omitempty"`
	// User assignment
	IdentificationStatus bool     `json:"IS"`
	IdentificationLevel  string   `json:"IL,omitempty"`
	IdentificationFlags  []string `json:"IF" validate:"max=4"`
	IdentificationType   string   `json:"IT"`
	IdentificationData   string   `json:"ID,omitempty"`
	TariffText           string   `json:"TT,omitempty"`
	// EVSE metrologic parameters
	LossCompensation LossCompensation `json:"LC,omitempty"`
	// Assignment of the charge point
	ChargePointIdentificationType string `json:"CT,omitempty"`
	ChargePointIdentification     string `json:"CI,omitempty"`
	// Readings
	Readings []Reading `json:"RD"`
}

type LossCompensation struct {
	Naming              string `json:"LN"`
	Identification      int    `json:"LI"`
	CableResistance     int    `json:"LR"`
	CableResistanceUnit int    `json:"LU"`
}

type Reading struct {
	Time              string `json:"TM"`
	Transaction       string `json:"TX,omitempty"`
	ReadingValue      int    `json:"RV"`
	ReadingIdentifier string `json:"RI,omitempty"`
	ReadingUnit       string `json:"RU"`
	ReadingType       string `json:"RT,omitempty"`
	CumulatedLoss     int    `json:"CL,omitempty"`
	ErrorFlags        string `json:"EF,omitempty"`
	Status            string `json:"ST"`
}

type MeterError string

const (
	MeterNotPresent   = MeterError("N")
	MeterOk           = MeterError("G")
	MeterTimeout      = MeterError("T")
	MeterDisconnected = MeterError("D")
	MeterRemoved      = MeterError("R")
	MeterManipulated  = MeterError("M")
	MeterExchanged    = MeterError("X")
	MeterIncompatible = MeterError("I")
	MeterOutOfRange   = MeterError("O")
	MeterSubstitute   = MeterError("S")
	MeterOtherError   = MeterError("E")
	MeterReadError    = MeterError("F")
)

type UserAssignmentState string

type ISO15118State string

const (
	ISO15118None          = ISO15118State("ISO15118_NONE")
	ISO15118PlugAndCharge = ISO15118State("ISO15118_PNC")
)

type OcppState string

const (
	OcppNone               = OcppState("OCPP_NONE")
	OcppRemoteStart        = OcppState("OCPP_RS")
	OcppAuthorizeMethod    = OcppState("OCPP_AUTH")
	OcppRemoteStartTLS     = OcppState("OCPP_RS_TLS")
	OcppAuthorizeMethodTLS = OcppState("OCPP_AUTH_TLS")
	OcppCache              = OcppState("OCPP_CACHE")
	OcppWhiteList          = OcppState("OCPP_WHITELIST")
	OcppCertified          = OcppState("OCPP_CERTIFIED")
)

type RfidState string

const (
	RfidNone         = RfidState("RFID_NONE")
	RfidPlain        = RfidState("RFID_PLAIN")
	RfidRelated      = RfidState("RFID_RELATED")
	RfidPreSharedKey = RfidState("RFID_PSK")
)

type ChargePointAssignmentType string

const (
	ChargePointAssignmentTypeEVSEID = ChargePointAssignmentType("EVSE_ID")
	ChargePointAssignmentTypeCBIDC  = ChargePointAssignmentType("CBIDC")
)

type TimeStatus string

const (
	TimeStatusUnknown      = TimeStatus("U")
	TimeStatusInformative  = TimeStatus("I")
	TimeStatusSynchronized = TimeStatus("S")
	TimeStatusRelative     = TimeStatus("R")
)

type Units string

const (
	UnitsWh       = Units("Wh")
	UnitskWh      = Units("kWh")
	UnitsMilliOhm = Units("mOhm")
	UnitsMicroOhm = Units("uOhm")
)

type CurrentType string

const (
	CurrentTypeAC = CurrentType("AC")
	CurrentTypeDC = CurrentType("DC")
)
