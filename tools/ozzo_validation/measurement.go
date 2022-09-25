package ozzo_validation

import (
	"deepfit/constants"
	validation "github.com/go-ozzo/ozzo-validation"
)

var Arm = []validation.Rule{
	validation.Max(200).Error(constants.EXCEEDED_MAX_VALUE),
	validation.Min(1).Error(constants.EXCEEDED_MIN_VALUE),
}

var Chest = []validation.Rule{
	validation.Max(300).Error(constants.EXCEEDED_MAX_VALUE),
	validation.Min(1).Error(constants.EXCEEDED_MIN_VALUE),
}

var Height = []validation.Rule{
	validation.Max(500).Error(constants.EXCEEDED_MAX_VALUE),
	validation.Min(1).Error(constants.EXCEEDED_MIN_VALUE),
}

var Hip = []validation.Rule{
	validation.Max(300).Error(constants.EXCEEDED_MAX_VALUE),
	validation.Min(1).Error(constants.EXCEEDED_MIN_VALUE),
}

var Leg = []validation.Rule{
	validation.Max(300).Error(constants.EXCEEDED_MAX_VALUE),
	validation.Min(1).Error(constants.EXCEEDED_MIN_VALUE),
}

var Shoulder = []validation.Rule{
	validation.Max(200).Error(constants.EXCEEDED_MAX_VALUE),
	validation.Min(1).Error(constants.EXCEEDED_MIN_VALUE),
}

var Waist = []validation.Rule{
	validation.Max(200).Error(constants.EXCEEDED_MAX_VALUE),
	validation.Min(1).Error(constants.EXCEEDED_MIN_VALUE),
}

var Weight = []validation.Rule{
	validation.Max(500).Error(constants.EXCEEDED_MAX_VALUE),
	validation.Min(1).Error(constants.EXCEEDED_MIN_VALUE),
}
