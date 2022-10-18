package dump

import "errors"

var Name = []string{
	"Kadir",
	"kadir",
	"Kadir123",
	"K",
	"Kadir DEniz DEneme Deneme Deneme Deneme",
	"Kadir DEniz DEneme Deneme Deneme Deneme /()!'",
	"^+'%&/(!",
	"Kadir !^^%&",
}

var NameResponses = []error{
	nil,
	nil,
	nil,
	nil,
	errors.New("wrong_length"),
	errors.New("wrong_length"),
	errors.New("wrong_format"),
	errors.New("wrong_format"),
}
