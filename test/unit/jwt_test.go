package unit

import (
	"deepfit/tools/jwt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestJWT(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "jwt")
}

var _ = Describe("Testing JWT Package", func() {

	var _jwt *jwt.JWT
	userId := primitive.NewObjectID()

	BeforeEach(func() {
		_jwt = jwt.New()
	})

	Context("Testing SetUserId", func() {
		It("should return set user id", func() {
			_jwt.SetUserId(userId)
			Expect(userId).To(Equal(_jwt.DumpClaim["userId"]))
		})
	})

	Context("Testing GetUserId", func() {
		It("should return get user id", func() {
			_jwt.DumpClaim["userId"] = userId
			Expect(userId).To(Equal(_jwt.GetUserId()))
		})
	})

	Context("Testing GetToken", func() {
		It("should return get token", func() {
			_jwt.DumpClaim["token"] = "token"
			Expect("token").To(Equal(_jwt.GetToken()))
		})
	})

	Context("Testing SetToken", func() {
		It("should return set token", func() {
			_jwt.SetToken("token")
			Expect("token").To(Equal(_jwt.DumpClaim["token"]))
		})
	})

	Context("Testing GetIsVerified", func() {
		It("should return get is verified", func() {
			_jwt.DumpClaim["is_verified"] = true
			Expect(true).To(Equal(_jwt.GetIsVerified()))
		})
	})

	Context("Testing SetIsVerified", func() {
		It("should return set is verified", func() {
			_jwt.SetIsVerified(true)
			Expect(true).To(Equal(_jwt.DumpClaim["is_verified"]))
		})
	})

	Context("Testing CreateToken", func() {
		It("should return create token", func() {
			_jwt.SetUserId(userId)
			_jwt.SetToken("token")
			_jwt.SetIsVerified(true)
			token := _jwt.CreateToken()
			Expect(token).ToNot(BeNil())
		})
	})

	Context("Testing DecodeToken", func() {
		It("should return decode token", func() {
			_jwt.SetUserId(userId).SetToken("token").SetIsVerified(true).CreateToken().DecodeToken()
			Expect(_jwt.GetUserId()).To(Equal(userId))
			Expect(_jwt.DumpClaim["token"]).To(Equal("token"))
			Expect(_jwt.DumpClaim["is_verified"]).To(Equal(true))
		})
	})
})
