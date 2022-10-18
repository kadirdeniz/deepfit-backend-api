package validation

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestValidationName(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "validation_name")
}

var _ = Describe("Testing ValidationName", func() {
	Context("Testing the correct value", func() {
		It("should return validate name", func() {
			Expect(validation.Validate("Kadir", ozzo_validation.Name...)).Should(BeNil())
			Expect(validation.Validate("kadir", ozzo_validation.Name...)).Should(BeNil())
			Expect(validation.Validate("K", ozzo_validation.Name...)).Should(BeNil())
		})
	})

	Context("Testing the wrong length", func() {
		It("should return validate name", func() {
			Expect(validation.Validate("Kadir DEniz DEneme Deneme Deneme Deneme", ozzo_validation.Name...).Error()).To(Equal("wrong_lenght"))
		})
	})

	// Testing Required Field
	Context("Testing the required field", func() {
		It("should return validate name", func() {
			Expect(validation.Validate("", ozzo_validation.Name...).Error()).To(Equal("required_field"))
		})
	})

	// Testing wrong type
	Context("Testing the wrong type", func() {
		It("should return validate name", func() {
			Expect(validation.Validate("123", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("!!!", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("'''", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("+++", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("%%%", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("&&&", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("(((", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("///", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate(")))", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("===", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("???", ozzo_validation.Name...).Error()).To(Equal("wrong_format"))
		})
	})
})
