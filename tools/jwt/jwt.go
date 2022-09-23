package jwt

import (
	"deepfit/configs"
	"fmt"
	"github.com/golang-jwt/jwt/v4"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return configs.JWT_SECRET, nil
}

type JWT struct {
	DumpClaim jwt.MapClaims
	Token     string
}

func New() *JWT {
	return &JWT{
		DumpClaim: jwt.MapClaims{
			"exp": configs.JWT_TTL,
		},
		Token: "",
	}
}

func (_jwt *JWT) SetUserId(userId primitive.ObjectID) *JWT {
	_jwt.DumpClaim["userId"] = userId
	return _jwt
}

func (_jwt *JWT) GetUserId() primitive.ObjectID {
	userId, _ := primitive.ObjectIDFromHex(_jwt.DumpClaim["userId"].(string))
	return userId
}

func (_jwt *JWT) GetToken() string {
	token := _jwt.DumpClaim["token"]
	return token.(string)
}

func (_jwt *JWT) SetToken(token string) *JWT {
	_jwt.DumpClaim["token"] = token
	return _jwt
}

func (_jwt *JWT) GetIsVerified() bool {
	isVerified := _jwt.DumpClaim["is_verified"]
	return isVerified.(bool)
}

func (_jwt *JWT) SetIsVerified(isVerified bool) *JWT {
	_jwt.DumpClaim["is_verified"] = isVerified
	return _jwt
}

func (_jwt *JWT) CreateToken() *JWT {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, _jwt.DumpClaim)
	tokenString, _ := token.SignedString([]byte(configs.JWT_SECRET))
	_jwt.SetToken(tokenString)
	return _jwt
}

func (_jwt *JWT) DecodeToken() *JWT {
	jwt.ParseWithClaims(_jwt.GetToken(), _jwt.DumpClaim, keyFunc)
	return _jwt
}
