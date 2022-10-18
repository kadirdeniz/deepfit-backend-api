package unit

import (
	"deepfit/pkg"
	"deepfit/test/dump"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "string")
}

var _ = Describe("Testing String Package", func() {
	Context("Testing UpperCaseFirstLetters", func() {
		It("should return upper case first letters", func() {
			for index, v := range dump.String {
				Expect(pkg.UpperCaseFirstLetters(v)).To(Equal(dump.UpperCaseFirstLetterResponses[index]))
			}
		})
	})

	Context("Testing LowerCaseReplaceEmpty", func() {
		It("should return lower case replace empty", func() {
			for index, v := range dump.String {
				Expect(pkg.LowerCaseReplaceEmpty(v)).To(Equal(dump.LowerCaseReplaceEmptyResponses[index]))
			}
		})
	})

	Context("Testing HashImageName", func() {
		It("should return hash image name", func() {
			for index, v := range dump.String {
				hashedImageName := pkg.HashImageName(v)
				if strings.Contains(hashedImageName, "-") {
					Expect(strings.Split(hashedImageName, "-")[1]).To(Equal(dump.HashImageNameResponses[index]))
				} else {
					Expect(hashedImageName).To(Equal(dump.HashImageNameResponses[index]))
				}
			}
		})
	})

	Context("Testing RandomCode", func() {
		It("should return random code", func() {
			randomCode := pkg.RandomCode()
			Expect(randomCode).To(BeNumerically(">", 10000))
			Expect(randomCode).To(BeNumerically("<", 99999))
		})
	})
})
