package ocmf_go

import "github.com/go-playground/validator/v10"

var messageValidator = validator.New()

func init() {
	// Register custom validators for the validator
	messageValidator.RegisterValidation("meterError", meterErrorValidator)
	messageValidator.RegisterValidation("userAssignmentState", userAssignmentStateValidator)
	messageValidator.RegisterValidation("pagination", paginationValidator)
	messageValidator.RegisterValidation("iso5118State", iso5118StateValidator)
	messageValidator.RegisterValidation("ocppState", ocppStateValidator)
	messageValidator.RegisterValidation("rfidState", rfidStateValidator)
	messageValidator.RegisterValidation("chargePointAssignment", chargePointAssignmentValidator)
	messageValidator.RegisterValidation("timeStatus", timeStatusValidatorValidator)
	messageValidator.RegisterValidation("time", timeValidator)
	messageValidator.RegisterValidation("unit", unitValidator)
	messageValidator.RegisterValidation("currentType", currentTypeValidator)

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
