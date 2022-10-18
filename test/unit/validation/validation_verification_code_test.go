package validation

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestValidationValidationCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "verification_code")
}

var _ = Describe("Testing VerificationCode", func() {
	Context("Testing the correct value", func() {
		It("should return validate verification_code", func() {
			Expect(validation.Validate(12345, ozzo_validation.VerificationCode...)).Should(BeNil())
		})
	})

	// Testing wrong length
	Context("Testing the wrong length", func() {
		It("should return validate verification_code", func() {
			Expect(validation.Validate(9999999, ozzo_validation.VerificationCode...).Error()).To(Equal("wrong_lenght"))
			Expect(validation.Validate(1, ozzo_validation.VerificationCode...).Error()).To(Equal("wrong_lenght"))
		})
	})

	// Testing Required Field
	Context("Testing the required field", func() {
		It("should return validate verification_code", func() {
			Expect(validation.Validate("", ozzo_validation.VerificationCode...).Error()).To(Equal("required_field"))
		})
	})
})
