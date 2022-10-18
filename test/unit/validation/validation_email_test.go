package validation

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestValidationEmail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "validation_email")
}

var _ = Describe("Testing ValidationEmail", func() {
	Context("Testing the correct value", func() {
		It("should return validate email", func() {
			Expect(validation.Validate("kadirdeniz@mail.com", ozzo_validation.Email...)).Should(BeNil())
			Expect(validation.Validate("kadirdeniz@gmail.com", ozzo_validation.Email...)).Should(BeNil())
			Expect(validation.Validate("kadirdeniz@hotmail.com", ozzo_validation.Email...)).Should(BeNil())
		})
	})

	Context("Testing the wrong length", func() {
		It("should return validate email", func() {
			Expect(validation.Validate("kadirdeniz@maildenememaildenemeuzunmailmaildenememaildenemeuzunmail.com", ozzo_validation.Email...).Error()).To(Equal("wrong_lenght"))
		})
	})

	// Testing Required Field
	Context("Testing the required field", func() {
		It("should return validate email", func() {
			Expect(validation.Validate("", ozzo_validation.Email...).Error()).To(Equal("required_field"))
		})
	})

	// Testing wrong type
	Context("Testing the wrong type", func() {
		It("should return validate email", func() {
			Expect(validation.Validate("kadir.deniz", ozzo_validation.Email...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("kadir.deniz@mail", ozzo_validation.Email...).Error()).To(Equal("wrong_format"))
			Expect(validation.Validate("kadir.deni  z@mail.com", ozzo_validation.Email...).Error()).To(Equal("wrong_format"))
		})
	})
})
