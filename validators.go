package ocmf_go

import "github.com/go-playground/validator/v10"

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
	must(messageValidator.RegisterValidation("time", timeValidator))
	must(messageValidator.RegisterValidation("unit", unitValidator))
	must(messageValidator.RegisterValidation("currentType", currentTypeValidator))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func meterErrorValidator(validator.FieldLevel) bool {
	return false
}

func userAssignmentStateValidator(validator.FieldLevel) bool {
	return false
}

func paginationValidator(validator.FieldLevel) bool {
	return false
}

func iso5118StateValidator(validator.FieldLevel) bool {
	return false
}

func ocppStateValidator(validator.FieldLevel) bool {
	return false
}

func rfidStateValidator(validator.FieldLevel) bool {
	return false
}

func chargePointAssignmentValidator(validator.FieldLevel) bool {
	return false
}

func timeStatusValidatorValidator(validator.FieldLevel) bool {
	return false
}

func timeValidator(validator.FieldLevel) bool {
	return false
}

func unitValidator(validator.FieldLevel) bool {
	return false
}

func currentTypeValidator(validator.FieldLevel) bool {
	return false
}
