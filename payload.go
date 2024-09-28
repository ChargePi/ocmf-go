package ocmf_go

const OcmfVersion = "0.4"

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

func isValidMeterError(me MeterError) bool {
	switch me {
	case MeterNotPresent, MeterOk, MeterTimeout, MeterDisconnected, MeterRemoved, MeterManipulated,
		MeterExchanged, MeterIncompatible, MeterOutOfRange, MeterSubstitute, MeterOtherError, MeterReadError:
		return true
	default:
		return false
	}
}

type UserAssignmentState string

const (
	UserAssignmentStateNONE      = UserAssignmentState("NONE")
	UserAssignmentStateHearsay   = UserAssignmentState("HEARSAY")
	UserAssignmentStateTrusted   = UserAssignmentState("TRUSTED")
	UserAssignmentStateVerified  = UserAssignmentState("VERIFIED")
	UserAssignmentStateCertified = UserAssignmentState("CERTIFIED")
	UserAssignmentStateSecure    = UserAssignmentState("SECURE")
	UserAssignmentStateMismatch  = UserAssignmentState("MISMATCH")
	UserAssignmentStateInvalid   = UserAssignmentState("INVALID")
	UserAssignmentStateOutdated  = UserAssignmentState("OUTDATED")
	UserAssignmentStateUnknown   = UserAssignmentState("UNKNOWN")
)

func isValidUserAssignmentState(state UserAssignmentState) bool {
	switch state {
	case UserAssignmentStateNONE, UserAssignmentStateHearsay, UserAssignmentStateTrusted,
		UserAssignmentStateVerified, UserAssignmentStateCertified, UserAssignmentStateSecure,
		UserAssignmentStateMismatch, UserAssignmentStateInvalid, UserAssignmentStateOutdated,
		UserAssignmentStateUnknown:
		return true
	default:
		return false
	}
}

type UserAssignmentType string

const ()

type ISO15118State string

const (
	ISO15118None          = ISO15118State("ISO15118_NONE")
	ISO15118PlugAndCharge = ISO15118State("ISO15118_PNC")
)

func isValidISO15118State(state ISO15118State) bool {
	switch state {
	case ISO15118None, ISO15118PlugAndCharge:
		return true
	default:
		return false
	}
}

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

func isValidOcppState(state OcppState) bool {
	switch state {
	case OcppNone, OcppRemoteStart, OcppAuthorizeMethod,
		OcppRemoteStartTLS, OcppAuthorizeMethodTLS,
		OcppCache, OcppWhiteList, OcppCertified:
		return true
	default:
		return false
	}
}

type RfidState string

const (
	RfidNone         = RfidState("RFID_NONE")
	RfidPlain        = RfidState("RFID_PLAIN")
	RfidRelated      = RfidState("RFID_RELATED")
	RfidPreSharedKey = RfidState("RFID_PSK")
)

func isValidRfidState(state RfidState) bool {
	switch state {
	case RfidNone, RfidPlain, RfidRelated, RfidPreSharedKey:
		return true
	default:
		return false
	}
}

type ChargePointAssignmentType string

const (
	ChargePointAssignmentTypeEVSEID = ChargePointAssignmentType("EVSE_ID")
	ChargePointAssignmentTypeCBIDC  = ChargePointAssignmentType("CBIDC")
)

func isValidChargePointAssignmentType(t ChargePointAssignmentType) bool {
	switch t {
	case ChargePointAssignmentTypeEVSEID, ChargePointAssignmentTypeCBIDC:
		return true
	default:
		return false
	}
}

type TimeStatus string

const (
	TimeStatusUnknown      = TimeStatus("U")
	TimeStatusInformative  = TimeStatus("I")
	TimeStatusSynchronized = TimeStatus("S")
	TimeStatusRelative     = TimeStatus("R")
)

func isValidTimeStatus(ts TimeStatus) bool {
	switch ts {
	case TimeStatusUnknown, TimeStatusInformative, TimeStatusSynchronized, TimeStatusRelative:
		return true
	default:
		return false
	}
}

type Units string

const (
	UnitsWh       = Units("Wh")
	UnitskWh      = Units("kWh")
	UnitsMilliOhm = Units("mOhm")
	UnitsMicroOhm = Units("uOhm")
)

func isValidUnit(u Units) bool {
	switch u {
	case UnitsWh, UnitskWh, UnitsMilliOhm, UnitsMicroOhm:
		return true
	default:
		return false
	}
}

type CurrentType string

const (
	CurrentTypeAC = CurrentType("AC")
	CurrentTypeDC = CurrentType("DC")
)

func isValidCurrentType(ct CurrentType) bool {
	switch ct {
	case CurrentTypeAC, CurrentTypeDC:
		return true
	default:
		return false
	}
}

type PayloadSection struct {
	// General information
	FormatVersion  string `json:"FV,omitempty"`
	GatewayID      string `json:"GI,omitempty"`
	GatewaySerial  string `json:"GS,omitempty"`
	GatewayVersion string `json:"GV,omitempty"`
	// Pagination
	Pagination string `json:"PG" validate:"required"`
	// Meter identification
	MeterVendor   string `json:"MV,omitempty"`
	MeterModel    string `json:"MM,omitempty"`
	MeterSerial   string `json:"MS" validate:"required"`
	MeterFirmware string `json:"MF,omitempty"`
	// User assignment
	IdentificationStatus bool     `json:"IS" validate:"required,"`
	IdentificationLevel  string   `json:"IL,omitempty" validate:"omitempty,userAssignmentState"`
	IdentificationFlags  []string `json:"IF" validate:"omitempty,max=4"`
	IdentificationType   string   `json:"IT" validate:"required,rfidState"`
	IdentificationData   string   `json:"ID,omitempty" validate:"omitempty,hex"`
	TariffText           string   `json:"TT,omitempty" validate:"omitempty,max=250"`
	// EVSE metrologic parameters
	LossCompensation LossCompensation `json:"LC,omitempty"`
	// Assignment of the charge point
	ChargeControllerVersion       string `json:"CF,omitempty" validate:"omitempty,max=25"`
	ChargePointIdentificationType string `json:"CT,omitempty" validate:"omitempty,chargePointAssignment"`
	ChargePointIdentification     string `json:"CI,omitempty"`
	// Readings
	Readings []Reading `json:"RD" validate:"required,dive"`
}

func (p *PayloadSection) Validate() error {
	return messageValidator.Struct(p)
}

type LossCompensation struct {
	Naming              string `json:"LN"`
	Identification      int    `json:"LI"`
	CableResistance     int    `json:"LR"`
	CableResistanceUnit int    `json:"LU"`
}

type Reading struct {
	Time              string `json:"TM" validate:"required,iso8601"`
	Transaction       string `json:"TX,omitempty" validate:"omitempty,oneof=B C X E L R A P S T"`
	ReadingValue      int    `json:"RV" validate:"required"`
	ReadingIdentifier string `json:"RI,omitempty"`
	ReadingUnit       string `json:"RU" validate:"required,unit"`
	ReadingType       string `json:"RT,omitempty" validate:"omitempty,currentType"`
	CumulatedLoss     int    `json:"CL,omitempty"`
	ErrorFlags        string `json:"EF,omitempty" validate:"omitempty,oneof=E t"`
	Status            string `json:"ST" validate:"required,meterError"`
}

func (r *Reading) Validate() error {
	return messageValidator.Struct(r)
}
