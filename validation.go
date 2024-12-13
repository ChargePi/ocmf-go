package ocmf_go

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var messageValidator = validator.New()

func init() {
	// Register custom validators for the validator
	must(messageValidator.RegisterValidation("meterError", meterErrorValidator))
	must(messageValidator.RegisterValidation("userAssignmentState", userAssignmentStateValidator))
	must(messageValidator.RegisterValidation("pagination", paginationValidator))
	must(messageValidator.RegisterValidation("iso5118State", iso5118StateValidator))
	must(messageValidator.RegisterValidation("ocppState", ocppStateValidator))
	must(messageValidator.RegisterValidation("rfidState", rfidStateValidator))
	must(messageValidator.RegisterValidation("chargePointAssignment", chargePointAssignmentValidator))
	must(messageValidator.RegisterValidation("timeStatus", timeStatusValidatorValidator))
	must(messageValidator.RegisterValidation("unit", unitValidator))
	must(messageValidator.RegisterValidation("currentType", currentTypeValidator))
	must(messageValidator.RegisterValidation("iso8601", iso8601WithMillisValidator))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func meterErrorValidator(fl validator.FieldLevel) bool {
	return isValidMeterError(MeterError(fl.Field().String()))
}

func userAssignmentStateValidator(fl validator.FieldLevel) bool {
	return isValidUserAssignmentState(UserAssignmentState(fl.Field().String()))
}

var indicatorNumberRegex = regexp.MustCompile(`^[TF]\d+$`)

func paginationValidator(fl validator.FieldLevel) bool {
	return indicatorNumberRegex.MatchString(fl.Field().String())
}

func iso5118StateValidator(fl validator.FieldLevel) bool {
	return isValidISO15118State(ISO15118State(fl.Field().String()))
}

func ocppStateValidator(fl validator.FieldLevel) bool {
	return isValidOcppState(OcppState(fl.Field().String()))
}

func rfidStateValidator(fl validator.FieldLevel) bool {
	return isValidRfidState(RfidState(fl.Field().String()))
}

func chargePointAssignmentValidator(fl validator.FieldLevel) bool {
	return isValidChargePointAssignmentType(ChargePointAssignmentType(fl.Field().String()))
}

func timeStatusValidatorValidator(fl validator.FieldLevel) bool {
	return isValidTimeStatus(TimeStatus(fl.Field().String()))
}

func unitValidator(fl validator.FieldLevel) bool {
	return isValidUnit(Units(fl.Field().String()))
}

func currentTypeValidator(fl validator.FieldLevel) bool {
	return isValidCurrentType(CurrentType(fl.Field().String()))
}

var iso8601WithMillisRegex = regexp.MustCompile(`^(?:19|20)\d{2}-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12]\d|3[01])T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:Z|-0[1-9]|-1\d|-2[0-3]|-00:?(?:0[1-9]|[1-5]\d)|\+[01]\d|\+2[0-3])(?:|:?[0-5]\d)$`)

func iso8601WithMillisValidator(fl validator.FieldLevel) bool {
	return iso8601WithMillisRegex.MatchString(fl.Field().String())
}

// Signature validation

var signatureValidator = validator.New()

func init() {
	// Register custom validators for the validator
	must(signatureValidator.RegisterValidation("signatureAlgorithm", signatureAlgorithmValidator))
	must(signatureValidator.RegisterValidation("signatureEncoding", signatureEncodingValidator))
}

func signatureEncodingValidator(fl validator.FieldLevel) bool {
	return isValidSignatureEncoding(SignatureEncoding(fl.Field().String()))
}

func signatureAlgorithmValidator(fl validator.FieldLevel) bool {
	return isValidSignatureAlgorithm(SignatureAlgorithm(fl.Field().String()))
}
