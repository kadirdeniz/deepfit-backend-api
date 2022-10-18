package unit

import (
	"deepfit/tools/bcrypt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBcrypt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "bcrypt")
}

var _ = Describe("Testing Bcrypt Package", func() {

	password := "password123!'^^"
	Context("Testing HashPassword", func() {
		It("should return hash password", func() {
			Expect(bcrypt.HashPassword(password)).ShouldNot(BeNil())
		})
	})

	Context("Testing ComparePasswords", func() {
		It("should return compare passwords", func() {
			hashedPassword := bcrypt.HashPassword(password)
			Expect(bcrypt.ComparePasswords(password, hashedPassword)).Should(BeTrue())
		})
	})

})
