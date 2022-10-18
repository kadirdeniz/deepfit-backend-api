package validation

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestValidationPassword(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "validation_password")
}

var _ = Describe("Testing ValidationEmail", func() {
	Context("Testing the correct value", func() {
		It("should return validate password", func() {
			Expect(validation.Validate("KadirDeniz", ozzo_validation.Password...)).Should(BeNil())
			Expect(validation.Validate("Kadir Deniz", ozzo_validation.Password...)).Should(BeNil())
		})
	})

	Context("Testing the wrong length", func() {
		It("should return validate password", func() {
			Expect(validation.Validate("Kadir DEniz DEneme Deneme Deneme Deneme", ozzo_validation.Password...).Error()).To(Equal("wrong_lenght"))
			Expect(validation.Validate("Kad", ozzo_validation.Password...).Error()).To(Equal("wrong_lenght"))
		})
	})

	// Testing Required Field
	Context("Testing the required field", func() {
		It("should return validate password", func() {
			Expect(validation.Validate("", ozzo_validation.Password...).Error()).To(Equal("required_field"))
		})
	})

})
