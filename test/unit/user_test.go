package unit

import (
	"deepfit/constants"
	"deepfit/internal/user"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "user")
}

var _ = Describe("Testing User Package", func() {

	// NewUser func should remove and set to integration test folder
	Context("Testing NewUser", func() {
		When("User testing with correct values", func() {
			It("should return new user", func() {
				userObj := user.NewUser("Kadir", "Deniz", "kadirdenz", "5061110088", "123456789")
				Expect(userObj.ID).ToNot(BeNil())
				Expect(userObj.Name).To(Equal("Kadir"))
				Expect(userObj.Surname).To(Equal("Deniz"))
				Expect(userObj.Nickname).To(Equal("kadirdenz"))
				Expect(userObj.Phone.Phone).To(Equal("5061110088"))
				Expect(userObj.Password).ToNot(Equal("123456789"))
				Expect(userObj.ProfilePhoto.GetImageName()).To(Equal(constants.DEFAULT_IMAGE))
				Expect(userObj.CoverPhoto.GetImageName()).To(Equal(constants.DEFAULT_IMAGE))
				Expect(userObj.Date.CreatedAt).ToNot(BeNil())
				Expect(userObj.Date.UpdatedAt).To(Equal(time.Time{}))
				Expect(userObj.Date.DeletedAt).To(Equal(time.Time{}))
			})
		})

		When("User testing string uppercase first letters for name and surname", func() {
			It("should return new user", func() {
				userObj := user.NewUser("kadir deniz", "deniz deneme", "kadirdenz", "5061110088", "123456789")
				Expect(userObj.Name).To(Equal("Kadir Deniz"))
				Expect(userObj.Surname).To(Equal("Deniz Deneme"))
				Expect(userObj.Date.CreatedAt).ToNot(BeNil())
				Expect(userObj.Date.UpdatedAt).To(Equal(time.Time{}))
				Expect(userObj.Date.DeletedAt).To(Equal(time.Time{}))
			})
		})

		When("Nickname is uppercase", func() {
			It("should return new user", func() {
				userObj := user.NewUser("kadir deniz", "deniz deneme", "KADIRDENIZ", "5061110088", "123456789")
				Expect(userObj.Nickname).To(Equal("kadirdeniz"))
				Expect(userObj.Date.CreatedAt).ToNot(BeNil())
				Expect(userObj.Date.UpdatedAt).To(Equal(time.Time{}))
				Expect(userObj.Date.DeletedAt).To(Equal(time.Time{}))
			})
		})

		When("Nickname is more than 1 word", func() {
			It("should return new user", func() {
				userObj := user.NewUser("kadir deniz", "deniz deneme", "kadir deniz", "5061110088", "123456789")
				Expect(userObj.Nickname).To(Equal("kadirdeniz"))
				Expect(userObj.Date.CreatedAt).ToNot(BeNil())
				Expect(userObj.Date.UpdatedAt).To(Equal(time.Time{}))
				Expect(userObj.Date.DeletedAt).To(Equal(time.Time{}))
			})
		})
	})

	Context("Testing SetID", func() {
		It("should return new user", func() {
			userObj := user.User{}
			userId := primitive.NewObjectID()
			userObj.SetId(userId)
			Expect(userObj.ID).To(Equal(userId))

			userObj.SetRandomId()
			Expect(userObj.ID).ToNot(BeNil())
		})
	})

	Context("Testing SetRandomId", func() {
		It("should return new user", func() {
			userObj := user.User{}

			userObj.SetRandomId()
			Expect(userObj.ID).ToNot(BeNil())
		})
	})

	Context("Testing SetName", func() {
		When("Name is lower case", func() {
			It("should upper case first letters lower case rest", func() {
				userObj := &user.User{}
				userObj.SetName("kadir deniz")
				Expect(userObj.Name).To(Equal("Kadir Deniz"))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})

		When("Name is upper case", func() {
			It("should upper case first letters lower case rest", func() {
				userObj := &user.User{}
				userObj.SetName("KADİR DENİZ")
				Expect(userObj.Name).To(Equal("Kadir Deniz"))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})

		When("Name with whitespaces", func() {
			It("should upper case first letters lower case rest", func() {
				userObj := &user.User{}
				userObj.SetName("            kadir                deniz               ")
				Expect(userObj.Name).To(Equal("Kadir Deniz"))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})
	})

	Context("Testing SetSurname", func() {
		When("Surname is lower case", func() {
			It("should upper case first letters lower case rest", func() {
				userObj := &user.User{}
				userObj.SetSurname("kadir deniz")
				Expect(userObj.Surname).To(Equal("Kadir Deniz"))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})

		When("Surname is upper case", func() {
			It("should upper case first letters lower case rest", func() {
				userObj := &user.User{}
				userObj.SetSurname("KADİR DENİZ")
				Expect(userObj.Surname).To(Equal("Kadir Deniz"))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})

		When("Surname with whitespaces", func() {
			It("should upper case first letters lower case rest", func() {
				userObj := &user.User{}
				userObj.SetSurname("            kadir                deniz               ")
				Expect(userObj.Surname).To(Equal("Kadir Deniz"))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})
	})

	Context("Testing SetNickname", func() {
		When("Nickname is upper case", func() {
			It("should lower case all and remove white spaces", func() {
				userObj := &user.User{}
				userObj.SetNickname("KADİR DENİZ")
				Expect(userObj.Nickname).To(Equal("kadirdeniz"))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})

		When("Nicknamne with whitespaces", func() {
			It("should lower case all and remove white spaces", func() {
				userObj := &user.User{}
				userObj.SetNickname(" kadir deniz ")
				Expect(userObj.Nickname).To(Equal("kadirdeniz"))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})
	})

	Context("Testing SetPhone", func() {
		It("should return set phone", func() {
			userObj := &user.User{}
			userObj.SetPhone("5061110088")
			Expect(userObj.Phone.Phone).To(Equal("5061110088"))
		})
	})

	Context("Testing SetPhoneObj", func() {
		It("should return set phone obj", func() {
			userObj := &user.User{}
			userObj.SetPhoneObj("5061110088")
			Expect(userObj.Phone.Phone).To(Equal("5061110088"))
			Expect(userObj.Phone.VerificationCode).ToNot(BeNil())
			Expect(userObj.Phone.IsVerified).To(BeFalse())
			Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
		})
	})

	Context("Testing SetEmail", func() {
		When("Email has upper case", func() {
			It("should return lower case and without white spaces email", func() {
				userObj := &user.User{}
				userObj.SetEmail("KADİR DENİZ@mail.com")
				Expect(userObj.Email.Email).To(Equal("kadirdeniz@mail.com"))
			})
		})

		When("Email has white spaces", func() {
			It("should return lower case and without white spaces email", func() {
				userObj := &user.User{}
				userObj.SetEmail("  kadir deniz @mail . com  ")
				Expect(userObj.Email.Email).To(Equal("kadirdeniz@mail.com"))
			})
		})
	})

	Context("Testing SetEmailObj", func() {
		It("should return set email obj", func() {
			userObj := &user.User{}
			userObj.SetEmailObj("kadirdeniz@mail.com")
			Expect(userObj.Email.Email).To(Equal("kadirdeniz@mail.com"))
			Expect(userObj.Email.VerificationCode).ToNot(BeNil())
			Expect(userObj.Email.IsVerified).To(BeFalse())
			Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
		})
	})

	Context("Testing HashPassword", func() {
		It("should return hash password", func() {
			userObj := &user.User{}
			userObj.HashPassword("123456")
			Expect(userObj.Password).ToNot(Equal("123456"))
			Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
		})
	})

	Context("Testing SetProfilePhoto", func() {
		When("Photo is nil", func() {
			It("should return default image", func() {
				userObj := &user.User{}
				userObj.SetProfilePhoto("")
				Expect(userObj.ProfilePhoto.GetImageName()).To(Equal(constants.DEFAULT_IMAGE))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})

		When("Photo is provided", func() {
			It("should return default image", func() {
				userObj := &user.User{}
				userObj.SetProfilePhoto("image.png")
				Expect(userObj.ProfilePhoto.GetImageName()).To(Equal("image.png"))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})
	})

	Context("Testing SetCoverPhoto", func() {
		When("Photo is nil", func() {
			It("should return default image", func() {
				userObj := &user.User{}
				userObj.SetCoverPhoto("")
				Expect(userObj.CoverPhoto.GetImageName()).To(Equal(constants.DEFAULT_IMAGE))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})

		When("Photo is provided", func() {
			It("should return default image", func() {
				userObj := &user.User{}
				userObj.SetCoverPhoto("image.png")
				Expect(userObj.CoverPhoto.GetImageName()).To(Equal("image.png"))
				Expect(userObj.Date.UpdatedAt).ToNot(BeNil())
			})
		})
	})
})
