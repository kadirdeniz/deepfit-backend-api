package validation

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestValidationNickname(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "validation_nickname")
}

var _ = Describe("Testing ValidationNickname", func() {
	Context("Testing the correct value", func() {
		It("should return validate nickname", func() {
			Expect(validation.Validate("Kadir", ozzo_validation.Nickname...)).Should(BeNil())
			Expect(validation.Validate("kadir", ozzo_validation.Nickname...)).Should(BeNil())
			Expect(validation.Validate("K", ozzo_validation.Nickname...)).Should(BeNil())
		})
	})

	Context("Testing the wrong length", func() {
		It("should return validate nickname", func() {
			Expect(validation.Validate("Kadir DEniz DEneme Deneme Deneme Deneme", ozzo_validation.Nickname...).Error()).To(Equal("wrong_lenght"))
		})
	})

	// Testing Required Field
	Context("Testing the required field", func() {
		It("should return validate nickname", func() {
			Expect(validation.Validate("", ozzo_validation.Nickname...).Error()).To(Equal("required_field"))
		})
	})

	// Testing wrong type
	Context("Testing the wrong type", func() {
		It("should return validate nickname", func() {
			Expect(validation.Validate("!!!", ozzo_validation.Nickname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("'''", ozzo_validation.Nickname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("+++", ozzo_validation.Nickname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("%%%", ozzo_validation.Nickname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("&&&", ozzo_validation.Nickname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("(((", ozzo_validation.Nickname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("///", ozzo_validation.Nickname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate(")))", ozzo_validation.Nickname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("===", ozzo_validation.Nickname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("???", ozzo_validation.Nickname...).Error()).To(Equal("wrong_format"))
		})
	})
})
