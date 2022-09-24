package pkg

import (
	"deepfit/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var space = regexp.MustCompile(`\s+`)
var AlphaOnlyRegex = `^[a-zA-ZğüşöçİĞÜŞÖÇ]+$`
var AlphaNumericRegex = `^[a-zA-ZğüşöçİĞÜŞÖÇ0-9._]+$`

func UpperCaseFirstLetters(input string) string {
	c := cases.Title(language.Turkish)
	return c.String(strings.ToLowerSpecial(unicode.TurkishCase, strings.TrimSpace(space.ReplaceAllString(input, " "))))
}

func LowerCaseReplaceEmpty(input string) string {
	return strings.ToLower(strings.ReplaceAll(input, " ", ""))
}

func HashImageName(imageName string) string {
	if imageName == constants.DEFAULT_IMAGE {
		return imageName
	}
	return primitive.NewObjectID().Hex() + "-" + strings.ToLower(strings.Join(strings.Split(imageName, " "), "_"))
}

func RandomCode() int {
	rangeLower := 10000
	rangeUpper := 99999
	return rangeLower + rand.Intn(rangeUpper-rangeLower+1)
}
