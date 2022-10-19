package pkg

type Phone struct {
	Phone            string `bson:"phone" json:"phone"`
	VerificationCode int    `bson:"verification_code" json:"-"`
	IsVerified       bool   `bson:"is_verified" json:"-"`
}

func NewPhone(phone string) Phone {
	return Phone{
		Phone:            phone,
		VerificationCode: RandomCode(),
		IsVerified:       false,
	}
}

func (phone *Phone) SetVerify() *Phone {
	phone.IsVerified = true
	return phone
}

func (phone *Phone) CheckVerificationCode(verificationCode int) bool {
	return phone.VerificationCode == verificationCode
}

func (phone *Phone) SetVerificationCode() *Phone {
	phone.VerificationCode = RandomCode()
	return phone
}
