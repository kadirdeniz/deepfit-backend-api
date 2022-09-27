package pkg

import (
	"deepfit/constants"
	"encoding/json"
)

func JSONEncoder(body []byte, dto interface{}) {
	encodeError := json.Unmarshal(body, &dto)
	if encodeError != nil {
		panic(
			NewError(constants.StatusBadRequest, constants.BAD_REQUEST, encodeError),
		)
	}
}

func JSONDecoder(body interface{}) []byte {
	marshaledJSON, decodeError := json.Marshal(body)
	if decodeError != nil {
		panic(
			NewError(constants.StatusBadRequest, constants.BAD_REQUEST, decodeError),
		)
	}

	return marshaledJSON
}
