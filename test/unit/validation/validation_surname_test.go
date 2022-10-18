package validation

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestValidationSurname(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "validation_surname")
}

var _ = Describe("Testing ValidationSurname", func() {
	Context("Testing the correct value", func() {
		It("should return validate surname", func() {
			Expect(validation.Validate("Kadir", ozzo_validation.Surname...)).Should(BeNil())
			Expect(validation.Validate("kadir", ozzo_validation.Surname...)).Should(BeNil())
			Expect(validation.Validate("K", ozzo_validation.Surname...)).Should(BeNil())
		})
	})

	Context("Testing the wrong length", func() {
		It("should return validate surname", func() {
			Expect(validation.Validate("Kadir DEniz DEneme Deneme Deneme Deneme", ozzo_validation.Surname...).Error()).To(Equal("wrong_lenght"))
		})
	})

	// Testing Required Field
	Context("Testing the required field", func() {
		It("should return validate surname", func() {
			Expect(validation.Validate("", ozzo_validation.Surname...).Error()).To(Equal("required_field"))
		})
	})

	// Testing wrong type
	Context("Testing the wrong type", func() {
		It("should return validate surname", func() {
			Expect(validation.Validate("123", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("!!!", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("'''", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("+++", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("%%%", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("&&&", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("(((", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("///", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate(")))", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("===", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("???", ozzo_validation.Surname...).Error()).To(Equal("wrong_format"))
		})
	})
})
