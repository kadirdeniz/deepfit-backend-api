package pkg

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	return primitive.NewObjectID().Hex() + "-" + strings.ToLower(strings.Join(strings.Split(imageName, " "), "_"))
}
