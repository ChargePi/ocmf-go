package ocmf_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayload_Validate(t *testing.T) {

}

func TestReading_Validate(t *testing.T) {

}

func Test_isValidOcppState(t *testing.T) {
	tests := []struct {
		name  string
		state OcppState
		want  bool
	}{
		{
			name:  "OCPP_NONE",
			state: OcppNone,
			want:  true,
		},
		{
			name:  "OCPP_RS",
			state: OcppRemoteStart,
			want:  true,
		},
		{
			name:  "OCPP_AUTH",
			state: OcppAuthorizeMethod,
			want:  true,
		},
		{
			name:  "OCPP_RS_TLS",
			state: OcppRemoteStartTLS,
			want:  true,
		},
		{
			name:  "OCPP_AUTH_TLS",
			state: OcppAuthorizeMethodTLS,
			want:  true,
		},
		{
			name:  "OCPP_CACHE",
			state: OcppCache,
			want:  true,
		},
		{
			name:  "OCPP_WHITELIST",
			state: OcppWhiteList,
			want:  true,
		},
		{
			name:  "OCPP_CERTIFIED",
			state: OcppCertified,
			want:  true,
		},
		{
			name:  "invalid",
			state: OcppState("invalid"),
			want:  false,
		},
		{
			name:  "Empty",
			state: OcppState(""),
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isValidOcppState(test.state)
			assert.Equal(t, test.want, res)
		})
	}
}

func Test_isValidRfidState(t *testing.T) {
	tests := []struct {
		name  string
		state RfidState
		want  bool
	}{
		{
			name:  "RFID_NONE",
			state: RfidNone,
			want:  true,
		},
		{
			name:  "RFID_PLAIN",
			state: RfidPlain,
			want:  true,
		},
		{
			name:  "RFID_RELATED",
			state: RfidRelated,
			want:  true,
		},
		{
			name:  "RFID_PSK",
			state: RfidPreSharedKey,
			want:  true,
		},
		{
			name:  "invalid",
			state: RfidState("invalid"),
			want:  false,
		},
		{
			name:  "Empty",
			state: RfidState(""),
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isValidRfidState(test.state)
			assert.Equal(t, test.want, res)
		})
	}
}

func Test_isValidChargePointAssignmentType(t *testing.T) {
	tests := []struct {
		name string
		t    ChargePointAssignmentType
		want bool
	}{
		{
			name: "EVSE_ID",
			t:    ChargePointAssignmentTypeEVSEID,
			want: true,
		},
		{
			name: "CBIDC",
			t:    ChargePointAssignmentTypeCBIDC,
			want: true,
		},
		{
			name: "invalid",
			t:    "invalid",
			want: false,
		},
		{
			name: "Empty",
			t:    "",
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isValidChargePointAssignmentType(test.t)
			assert.Equal(t, test.want, res)
		})
	}
}

func Test_isValidISO15118State(t *testing.T) {
	tests := []struct {
		name  string
		state ISO15118State
		want  bool
	}{
		{
			name:  "ISO15118None",
			state: ISO15118None,
			want:  true,
		},
		{
			name:  "ISO15118PlugAndCharge",
			state: ISO15118PlugAndCharge,
			want:  true,
		},
		{
			name:  "invalid",
			state: "invalid",
			want:  false,
		},
		{
			name:  "Empty",
			state: "",
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isValidISO15118State(test.state)
			assert.Equal(t, test.want, res)
		})
	}
}

func Test_isValidMeterError(t *testing.T) {
	tests := []struct {
		name  string
		state MeterError
		want  bool
	}{
		{
			name:  "MeterNotPresent",
			state: MeterNotPresent,
			want:  true,
		},
		{
			name:  "MeterOk",
			state: MeterOk,
			want:  true,
		},
		{
			name:  "MeterTimeout",
			state: MeterTimeout,
			want:  true,
		},
		{
			name:  "MeterDisconnected",
			state: MeterDisconnected,
			want:  true,
		},
		{
			name:  "MeterRemoved",
			state: MeterRemoved,
			want:  true,
		},
		{
			name:  "MeterManipulated",
			state: MeterManipulated,
			want:  true,
		},
		{
			name:  "MeterExchanged",
			state: MeterExchanged,
			want:  true,
		},
		{
			name:  "MeterIncompatible",
			state: MeterIncompatible,
			want:  true,
		},
		{
			name:  "MeterOutOfRange",
			state: MeterOutOfRange,
			want:  true,
		},
		{
			name:  "MeterSubstitute",
			state: MeterSubstitute,
			want:  true,
		},
		{
			name:  "MeterOtherError",
			state: MeterOtherError,
			want:  true,
		},
		{
			name:  "MeterReadError",
			state: MeterReadError,
			want:  true,
		},
		{
			name:  "invalid",
			state: MeterError("invalid"),
			want:  false,
		},
		{
			name:  "Empty",
			state: MeterError(""),
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isValidMeterError(test.state)
			assert.Equal(t, test.want, res)
		})
	}
}

func Test_isValidTimeStatus(t *testing.T) {
	tests := []struct {
		name string
		ts   TimeStatus
		want bool
	}{
		{
			name: "Unknown",
			ts:   TimeStatusUnknown,
			want: true,
		},
		{
			name: "TimeStatusSynchronized",
			ts:   TimeStatusSynchronized,
			want: true,
		},
		{
			name: "TimeStatusRelative",
			ts:   TimeStatusRelative,
			want: true,
		},
		{
			name: "TimeStatusInformative",
			ts:   TimeStatusInformative,
			want: true,
		},
		{
			name: "invalid",
			ts:   TimeStatus("invalid"),
			want: false,
		},
		{
			name: "Empty",
			ts:   "",
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isValidTimeStatus(test.ts)
			assert.Equal(t, test.want, res)
		})
	}
}

func Test_isValidUnit(t *testing.T) {
	tests := []struct {
		name string
		unit Units
		want bool
	}{
		{
			name: "Wh",
			unit: UnitsWh,
			want: true,
		},
		{
			name: "kWh",
			unit: UnitskWh,
			want: true,
		},
		{
			name: "mOhm",
			unit: UnitsMilliOhm,
			want: true,
		},
		{
			name: "uOhm",
			unit: UnitsMicroOhm,
			want: true,
		},
		{
			name: "invalid",
			unit: Units("invalid"),
			want: false,
		},
		{
			name: "Empty",
			unit: Units(""),
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isValidUnit(test.unit)
			assert.Equal(t, test.want, res)
		})
	}
}

func Test_isValidCurrentType(t *testing.T) {
	tests := []struct {
		name string
		ct   CurrentType
		want bool
	}{
		{
			name: "AC",
			ct:   CurrentTypeAC,
			want: true,
		},
		{
			name: "DC",
			ct:   CurrentTypeDC,
			want: true,
		},
		{
			name: "invalid",
			ct:   CurrentType("invalid"),
			want: false,
		},
		{
			name: "Empty",
			ct:   CurrentType(""),
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isValidCurrentType(test.ct)
			assert.Equal(t, test.want, res)
		})
	}
}
