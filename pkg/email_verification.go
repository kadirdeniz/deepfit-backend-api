package pkg

type Email struct {
	Email            string `bson:"email" json:"email"`
	VerificationCode int    `bson:"verification_code" json:"-"`
	IsVerified       bool   `bson:"is_verified" json:"-"`
}

func NewEmail(email string) Email {
	return Email{
		Email:            LowerCaseReplaceEmpty(email),
		VerificationCode: RandomCode(),
		IsVerified:       false,
	}
}

func (email *Email) SetVerify() *Email {
	email.IsVerified = true
	return email
}

func (email *Email) SetVerificationCode() *Email {
	email.VerificationCode = RandomCode()
	return email
}
